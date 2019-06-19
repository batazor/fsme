import React, { Component } from 'react'
import { withStyles } from '@material-ui/core/styles'
import _ from 'lodash'
import * as SRD from "storm-react-diagrams";

import "storm-react-diagrams/dist/style.min.css";

import TrayWidget from "./TrayWidget";
import TrayItemWidget from "./TrayItemWidget";

const styles = {
  root: {
    display: 'flex',
    flex: 1,
  },
  diagramLayer: {
    display: "flex",
    flex: 1,
    position: "relative",
    flexGrow: 1,
  },
  srd: {
    height: "100%",
    backgroundColor: "#3c3c3c !important",
    backgroundImage: "linear-gradient(0deg, transparent 24%, rgba(255, 255, 255, .05) 25%, rgba(255, 255, 255, .05) 26%, transparent 27%, transparent 74%, rgba(255, 255, 255, .05) 75%, rgba(255, 255, 255, .05) 76%, transparent 77%, transparent), linear-gradient(90deg, transparent 24%, rgba(255, 255, 255, .05) 25%, rgba(255, 255, 255, .05) 26%, transparent 27%, transparent 74%, rgba(255, 255, 255, .05) 75%, rgba(255, 255, 255, .05) 76%, transparent 77%, transparent)",
    backgroundSize: "50px 50px",
  },
};

class Graph extends Component {

  static getDerivedStateFromProps(props, state) {
    // if update state
    if (_.get(props, 'fsm.FSM.State') !== _.get(state, 'fsm.FSM.State')) {
      const engine = state.engine
      const mapNode = {}
      const edges = []
      let nodes = []

      var model = new SRD.DiagramModel();

      const Transitions = _.get(props, 'fsm.FSM.Transitions', null)
      if (Transitions) {
        Object.keys(Transitions).forEach((item, index) => mapNode[item] = index)

        Object.keys(Transitions).forEach(item => {
          const node = new SRD.DefaultNodeModel(item, "rgb(0,192,255)");

          // set UI
          const UI = _.get(props, `fsm.UI[${item}]`, {})
          const xPosition = UI.X || 100
          const yPosition = UI.Y || 100
          node.setPosition(xPosition, yPosition);

          nodes.push(node)
        })

        Object.keys(Transitions).forEach((item, index) => {
          const indexStartNode = mapNode[item]
          const startNode = nodes[indexStartNode]

          Object.keys(Transitions[item]).forEach(edge => {
            const indexEndNode = mapNode[edge]
            const endNode = nodes[indexEndNode]

            if (endNode) {
              const portOut = startNode.addOutPort("Out");
              const portIn = endNode.addInPort("In");

              const link = portOut.link(portIn);
              edges.push(link)
            }
          })
        })
      }

      model.addAll(...nodes, ...edges);

      engine.setDiagramModel(model);
      return { engine, fsm: props.fsm }
    }

    return null;
  }

  constructor() {
    super();

    // setup the diagram engine
    const engine = new SRD.DiagramEngine();
    engine.installDefaultFactories();

    // setup the diagram model
    const model = new SRD.DiagramModel();

    // load model into engine
    engine.setDiagramModel(model);

    this.state = {
      engine,
    }

    this.getDiagramEngine = this.getDiagramEngine.bind(this)
    this.onChangeFSM = this.onChangeFSM.bind(this)
  }

	getDiagramEngine() {
    this.onChangeFSM()
		return this.state.engine;
	}

  getNodeNameById(idNode) {
    const { nodes } = this.state.engine.diagramModel.serializeDiagram()
    const node = nodes.filter(item => item.id === idNode)
    if (node.length) {
      return node[0].name
    }

    return null
  }

  onChangeFSM() {
    console.warn('onChangeFSM', this.state.engine.diagramModel.serializeDiagram())
    console.warn('FSM', this.props.fsm.FSM)

    const { nodes, links } = this.state.engine.diagramModel.serializeDiagram()

    // Update Node
    nodes.forEach(node => {
      const isExistNode = _.get(this.state.fsm, `FSM.Transitions[${node.name}]`, {})
      _.set(this.state.fsm, `FSM.Transitions[${node.name}]`, isExistNode)

      // Update UI
      const nodeUI = _.get(this.state.fsm, `UI[${node.name}]`, {})
      nodeUI.X = node.x
      nodeUI.Y = node.y
      _.set(this.state.fsm, `UI[${node.name}]`, nodeUI)
    })

    // Update Link
    links.forEach(link => {
      const StartNode = this.getNodeNameById(link.source)
      const EndNode = this.getNodeNameById(link.target)

      if (EndNode) {
        const isExistNode = _.get(this.state.fsm, `FSM.Transitions[${StartNode}]`, {})
        isExistNode[EndNode] = {}
        _.set(this.state.fsm, `FSM.Transitions[${StartNode}]`, isExistNode)
      }
    })

    // const Transitions = _.get(props, 'fsm.FSM.Transitions', [])
    console.warn('STATE', this.state.fsm.FSM)

    this.props.onChange(this.state.fsm)
  }

  render() {
    const { classes } = this.props;

    return (
      <div className={classes.root}>
        <TrayWidget>
          <TrayItemWidget model={{ type: "in" }} name="In Node" color="rgb(192,255,0)" />
          <TrayItemWidget model={{ type: "out" }} name="Out Node" color="rgb(0,192,255)" />
        </TrayWidget>

        <div
					className={classes.diagramLayer}
					onDrop={event => {
						var data = JSON.parse(event.dataTransfer.getData("storm-diagram-node"));
						var nodesCount = _.keys(
							this
								.getDiagramEngine()
								.getDiagramModel()
								.getNodes()
						).length;

						var node = null;
						if (data.type === "in") {
							node = new SRD.DefaultNodeModel("Node " + (nodesCount + 1), "rgb(192,255,0)");
							node.addInPort("In");
						} else {
							node = new SRD.DefaultNodeModel("Node " + (nodesCount + 1), "rgb(0,192,255)");
							node.addOutPort("Out");
						}
						var points = this.getDiagramEngine().getRelativeMousePoint(event);
						node.x = points.x;
						node.y = points.y;
						this
							.getDiagramEngine()
							.getDiagramModel()
							.addNode(node);
						this.forceUpdate();
					}}
					onDragOver={event => {
						event.preventDefault();
					}}
          onMouseUp={event => this.onChangeFSM(event)}
				>
          <SRD.DiagramWidget className={classes.srd} diagramEngine={this.state.engine} />
        </div>
      </div>
    );
  }
}

export default withStyles(styles)(Graph);
