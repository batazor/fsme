import React, { Component } from 'react';
import { withStyles } from '@material-ui/core/styles';
import {
  GraphView, // required
} from 'react-digraph';

const GraphConfig =  {
  NodeTypes: {
    empty: { // required to show empty nodes
      typeText: "None",
      shapeId: "#empty", // relates to the type property of a node
      shape: (
        <symbol viewBox="0 0 100 100" id="empty" key="0">
          <circle cx="50" cy="50" r="45"></circle>
        </symbol>
      )
    },
    custom: { // required to show empty nodes
      typeText: "Custom",
      shapeId: "#custom", // relates to the type property of a node
      shape: (
        <symbol viewBox="0 0 100 100" id="custom" key="0">
          <circle style={{ fill: 'yellow' }} cx="50" cy="50" r="45"></circle>
        </symbol>
      )
    }
  },
  NodeSubtypes: {},
  EdgeTypes: {
    emptyEdge: {  // required to show empty edges
      shapeId: "#emptyEdge",
      shape: (
        <symbol viewBox="0 0 50 50" id="emptyEdge" key="0">
          <circle cx="25" cy="25" r="8" fill="currentColor"> </circle>
        </symbol>
      )
    }
  }
}

const NODE_KEY = "id"       // Allows D3 to correctly update DOM

const styles = {
  root: {
    display: 'flex',
    flex: 1,
  },
};

class Graph extends Component {

  constructor(props) {
    super(props);

    // console.log('fsm', props.fsm.FSM)
    const mapNode = {}
    Object.keys(props.fsm.FSM.Transitions).forEach((item, index) => mapNode[item] = index + 1)
    const edges = []

    Object.keys(props.fsm.FSM.Transitions).forEach(item => {
      Object.keys(props.fsm.FSM.Transitions[item]).forEach(edge => {
        edges.push({
          "source": mapNode[item],
          "target": mapNode[edge],
        })
      })
    })

    this.state = {
      "graph": {
        "nodes": Object.keys(props.fsm.FSM.Transitions).map((item, index) => ({
          "id": index + 1,
          "title": item,
          "x": 150 + Math.random() * index * 300 + (index % 2 ? 0 : 200),
          "y": 230 + Math.random() * index * 200 + (index % 2 ? 0 : 450),
          "type": item === props.fsm.FSM.State ? "custom" : "empty",
        })),
        "edges": edges,
      },
      "selected": {}
    }
  }

  /* Define custom graph editing methods here */

  render() {
    const nodes = this.state.graph.nodes;
    const edges = this.state.graph.edges;
    const selected = this.state.selected;

    const NodeTypes = GraphConfig.NodeTypes;
    const NodeSubtypes = GraphConfig.NodeSubtypes;
    const EdgeTypes = GraphConfig.EdgeTypes;

    const { classes } = this.props;

    return (
      <div className={classes.root}>
        <GraphView
          className={classes.root}
          ref='GraphView'
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
    );
  }
}


export default withStyles(styles)(Graph);
