import React, { Component } from 'react'
import { withStyles } from '@material-ui/core/styles'
import {
  GraphView, // required
} from 'react-digraph'
import _ from 'lodash'

const GraphConfig = {
  NodeTypes: {
    empty: { // required to show empty nodes
      typeText: 'None',
      shapeId: '#empty', // relates to the type property of a node
      shape: (
        <symbol viewBox="0 0 100 100" id="empty" key="0">
          <circle cx="50" cy="50" r="45" />
        </symbol>
      ),
    },
    custom: { // required to show empty nodes
      typeText: 'Custom',
      shapeId: '#custom', // relates to the type property of a node
      shape: (
        <symbol viewBox="0 0 100 100" id="custom" key="0">
          <circle style={{ fill: 'yellow' }} cx="50" cy="50" r="45" />
        </symbol>
      ),
    },
  },
  NodeSubtypes: {},
  EdgeTypes: {
    emptyEdge: { // required to show empty edges
      shapeId: '#emptyEdge',
      shape: (
        <symbol viewBox="0 0 50 50" id="emptyEdge" key="0">
          <circle cx="25" cy="25" r="8" fill="currentColor"> </circle>
        </symbol>
      ),
    },
  },
}

const NODE_KEY = 'id' // Allows D3 to correctly update DOM

const styles = {
  root: {
    display: 'flex',
    flex: 1,
  },
}

class Graph extends Component {
  static getDerivedStateFromProps(props, state) {
    // if update state
    if (_.get(props, 'fsm.FSM.State') !== _.get(state, 'fsm.FSM.State')) {
      const mapNode = {}
      const edges = []
      const nodes = []

      const Transitions = _.get(props, 'fsm.FSM.Transitions', null)
      if (Transitions) {
        Object.keys(Transitions).forEach((item, index) => mapNode[item] = index + 1)

        Object.keys(Transitions).forEach((item, index) => {
          Object.keys(Transitions[item]).forEach(edge => {
            edges.push({
              source: mapNode[item],
              target: mapNode[edge],
            })
          })

          // set UI
          const UI = _.get(props, `fsm.UI[${item}]`, {})
          const xPosition = UI.X || 100
          const yPosition = UI.Y || 100

          nodes.push({
            id: index + 1,
            title: item,
            x: xPosition,
            y: yPosition,
            type: item === props.fsm.FSM.State ? 'custom' : 'empty',
          })
        })
      }

      return {
        graph: {
          nodes,
          edges,
        },
        selected: {},
      }
    }

    return null
  }

  constructor() {
    super()

    this.state = {
      graph: {
        nodes: [],
        edges: [],
      },
      selected: {},
    }
  }

  /* Define custom graph editing methods here */

  render() {
    const { classes } = this.props // eslint-disable-line
    const { selected, graph } = this.state
    const { nodes, edges } = graph

    const { NodeTypes } = GraphConfig
    const { NodeSubtypes } = GraphConfig
    const { EdgeTypes } = GraphConfig

    return (
      <div className={classes.root}>
        <GraphView
          className={classes.root}
          ref="GraphView" // eslint-disable-line
          gridSize
          nodeKey={NODE_KEY}
          nodes={nodes}
          edges={edges}
          selected={selected}
          nodeTypes={NodeTypes}
          nodeSubtypes={NodeSubtypes}
          edgeTypes={EdgeTypes}
          onSelectNode={this.onSelectNode}
          onCreateNode={this.onCreateNode}
          onUpdateNode={this.onUpdateNode}
          onDeleteNode={this.onDeleteNode}
          onSelectEdge={this.onSelectEdge}
          onCreateEdge={this.onCreateEdge}
          onSwapEdge={this.onSwapEdge}
          onDeleteEdge={this.onDeleteEdge}
        />
      </div>
    )
  }
}


export default withStyles(styles)(Graph)
