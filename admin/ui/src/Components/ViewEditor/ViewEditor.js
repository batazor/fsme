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

    }

    return null;
  }

  constructor() {
    super();

    this.state = {}

    this.diagramEngine = new SRD.DiagramEngine();
		this.diagramEngine.installDefaultFactories();

		this.newModel();
  }

  /* Define custom graph editing methods here */
  newModel() {
		this.activeModel = new SRD.DiagramModel();
		this.diagramEngine.setDiagramModel(this.activeModel);

		//3-A) create a default node
		var node1 = new SRD.DefaultNodeModel("Node 1", "rgb(0,192,255)");
		let port = node1.addOutPort("Out");
		node1.setPosition(100, 100);

		//3-B) create another default node
		var node2 = new SRD.DefaultNodeModel("Node 2", "rgb(192,255,0)");
		let port2 = node2.addInPort("In");
		node2.setPosition(400, 100);

		// link the ports
		let link1 = port.link(port2);

		this.activeModel.addAll(node1, node2, link1);
	}

	getActiveDiagram(): SRD.DiagramModel {
		return this.activeModel;
	}

	getDiagramEngine(): SRD.DiagramEngine {
		return this.diagramEngine;
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
				>
          <SRD.DiagramWidget className={classes.srd} diagramEngine={this.getDiagramEngine()} />
        </div>
      </div>
    );
  }
}


export default withStyles(styles)(Graph);
