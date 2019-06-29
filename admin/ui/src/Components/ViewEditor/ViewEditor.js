import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'
import _ from 'lodash'
import * as SRD from 'storm-react-diagrams'

import 'storm-react-diagrams/dist/style.min.css'

import TrayWidget from './TrayWidget'
import TrayItemWidget from './TrayItemWidget'

const styles = {
  root: {
    display: 'flex',
    flex: 1,
  },
  diagramLayer: {
    display: 'flex',
    flex: 1,
    position: 'relative',
    flexGrow: 1,
  },
  srd: {
    height: '100%',
    backgroundColor: '#3c3c3c !important',
    backgroundImage: 'linear-gradient(0deg, transparent 24%, rgba(255, 255, 255, .05) 25%, rgba(255, 255, 255, .05) 26%, transparent 27%, transparent 74%, rgba(255, 255, 255, .05) 75%, rgba(255, 255, 255, .05) 76%, transparent 77%, transparent), linear-gradient(90deg, transparent 24%, rgba(255, 255, 255, .05) 25%, rgba(255, 255, 255, .05) 26%, transparent 27%, transparent 74%, rgba(255, 255, 255, .05) 75%, rgba(255, 255, 255, .05) 76%, transparent 77%, transparent)',
    backgroundSize: '50px 50px',
  },
}

class Graph extends Component {
  static getDerivedStateFromProps(props, state) {
    // if update state
    if (_.get(props, 'fsm.FSM.State') !== _.get(state, 'fsm.FSM.State')) {
      const { engine } = state
      const mapNode = {}
      const edges = []
      const nodes = []
      const ports = {}

      const model = new SRD.DiagramModel()

      const Transitions = _.get(props, 'fsm.FSM.Transitions', null)
      if (Transitions) {
        Object.keys(Transitions).forEach((item, index) => mapNode[item] = index)

        // Add node
        Object.keys(Transitions).forEach(item => {
          const color = props.fsm.FSM.State === item ? 'rgb(192,255,0)' : 'rgb(0,192,255)'
          const node = new SRD.DefaultNodeModel(item, color)

          // set UI
          const UI = _.get(props, `fsm.UI[${item}]`, {})
          const xPosition = UI.X || 100
          const yPosition = UI.Y || 100
          node.setPosition(xPosition, yPosition)

          // set Port
          ports[item] = {}
          switch (UI.TYPE) {
            case 'In':
              ports[item].In = node.addInPort('In')
              break
            case 'Out':
              ports[item].Out = node.addOutPort('Out')
              break
            case 'InOut':
              ports[item].Out = node.addOutPort('Out')
              ports[item].In = node.addInPort('In')
              break
            default:
          }

          nodes.push(node)
        })

        // Add link
        Object.keys(Transitions).forEach(item => {
          Object.keys(Transitions[item]).forEach(edge => {
            const indexEndNode = mapNode[edge]
            const endNode = nodes[indexEndNode]

            if (endNode) {
              const portOut = ports[edge].Out
              const portIn = ports[item].In

              if (portIn && portOut) {
                const link = portOut.link(portIn)

                // Add label [event]
                const Event = _.get(props, `fsm.FSM.Events[${edge}]`)
                if (Event) {
                  link.addLabel(Event)
                }

                edges.push(link)
              }
            }
          })
        })
      }

      model.addAll(...nodes, ...edges)

      engine.setDiagramModel(model)
      return { engine, fsm: props.fsm }
    }

    return null
  }

  constructor() {
    super()

    // setup the diagram engine
    const engine = new SRD.DiagramEngine()
    engine.installDefaultFactories()

    // setup the diagram model
    const model = new SRD.DiagramModel()

    // load model into engine
    engine.setDiagramModel(model)

    this.state = {
      engine,
    }

    this.getDiagramEngine = this.getDiagramEngine.bind(this)
    this.onChangeFSM = this.onChangeFSM.bind(this)
  }

  onChangeFSM() {
    const { onChange } = this.props
    const { fsm, engine } = this.state

    // /* eslint-disable */
    const { nodes, links } = engine.diagramModel.serializeDiagram()

    // Update Node
    nodes.forEach(node => {
      const isExistNode = _.get(fsm, `FSM.Transitions[${node.name}]`, {})
      _.set(fsm, `FSM.Transitions[${node.name}]`, isExistNode)

      // Update UI
      const nodeUI = _.get(fsm, `UI[${node.name}]`, {})
      nodeUI.X = parseInt(node.x, 10)
      nodeUI.Y = parseInt(node.y, 10)

      _.set(fsm, `UI[${node.name}]`, nodeUI)
    })

    // Update Link
    links.forEach(link => {
      const StartNode = this.getNodeNameById(link.target)
      const EndNode = this.getNodeNameById(link.source)

      if (StartNode && EndNode) {
        _.set(fsm, `FSM.Transitions[${StartNode}][${EndNode}]`, {})

        // Add event
        _.set(fsm, `FSM.Events[${EndNode}]`, EndNode)
      }
    })

    // /* eslint-enable */
    onChange(fsm)
  }

  onDrop(event) {
    const { fsm } = this.state
    const { onChange } = this.props
    const nodeUI = {}
    const data = JSON.parse(event.dataTransfer.getData('storm-diagram-node'))
    const nodesCount = _.keys(
      this
        .getDiagramEngine()
        .getDiagramModel()
        .getNodes(),
    ).length

    const nameNode = `Node ${nodesCount}`

    const node = new SRD.DefaultNodeModel(nameNode, 'rgb(192,255,0)')
    if (data.type === 'in') {
      node.addInPort('In')
      nodeUI.TYPE = 'In'
    }
    if (data.type === 'out') {
      nodeUI.TYPE = 'Out'
    }
    if (data.type === 'inOut') {
      nodeUI.TYPE = 'InOut'
    }

    const points = this.getDiagramEngine().getRelativeMousePoint(event)
    nodeUI.X = points.x
    nodeUI.Y = points.y

    _.set(fsm, `UI[${nameNode}]`, nodeUI)
    _.set(fsm, `FSM.Transitions[${nameNode}]`, {})

    onChange(fsm)
  }

  getDiagramEngine() {
    this.onChangeFSM()
    return this.state.engine // eslint-disable-line
  }

  getNodeNameById(idNode) {
    const { nodes } = this.state.engine.diagramModel.serializeDiagram() // eslint-disable-line
    const node = nodes.filter(item => item.id === idNode)
    if (node.length) {
      return node[0].name
    }

    return null
  }

  render() {
    const { classes } = this.props
    const { engine } = this.state

    return (
      <div className={classes.root}>
        <TrayWidget>
          <TrayItemWidget model={{ type: 'in' }} name="Start" color="rgb(192,255,0)" />
          <TrayItemWidget model={{ type: 'inOut' }} name="Event" color="rgb(192,255,0)" />
          <TrayItemWidget model={{ type: 'out' }} name="End" color="rgb(0,192,255)" />
        </TrayWidget>

        <div // eslint-disable-line
          className={classes.diagramLayer}
          onDrop={event => this.onDrop(event)}
          onDragOver={event => event.preventDefault()}
          onMouseUp={event => this.onChangeFSM(event)}
        >
          <SRD.DiagramWidget className={classes.srd} diagramEngine={engine} />
        </div>
      </div>
    )
  }
}

Graph.propTypes = {
  classes: PropTypes.object.isRequired, // eslint-disable-line

  onChange: PropTypes.func.isRequired,
}

export default withStyles(styles)(Graph)
