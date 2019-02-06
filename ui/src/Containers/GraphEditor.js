import React, { Component } from 'react';
import { withStyles } from '@material-ui/core/styles';
import {
  GraphView, // required
  Edge, // optional
  type IEdge, // optional
  Node, // optional
  type INode, // optional
  type LayoutEngineType, // required to change the layoutEngineType, otherwise optional
  BwdlTransformer, // optional, Example JSON transformer
  GraphUtils // optional, useful utility functions
} from 'react-digraph';

const sample = {
  "nodes": [
    {
      "id": 1,
      "title": "Node A",
      "x": 258.3976135253906,
      "y": 331.9783248901367,
      "type": "empty"
    },
    {
      "id": 2,
      "title": "Node B",
      "x": 593.9393920898438,
      "y": 260.6060791015625,
      "type": "empty"
    },
    {
      "id": 3,
      "title": "Node C",
      "x": 237.5757598876953,
      "y": 61.81818389892578,
      "type": "custom"
    },
    {
      "id": 4,
      "title": "Node C",
      "x": 600.5757598876953,
      "y": 600.81818389892578,
      "type": "custom"
    }
  ],
  "edges": [
    {
      "source": 1,
      "target": 2,
      "type": "emptyEdge"
    },
    {
      "source": 2,
      "target": 4,
      "type": "emptyEdge"
    }
  ]
}

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
        <symbol viewBox="0 0 50 25" id="custom" key="0">
          <ellipse cx="50" cy="25" rx="50" ry="25"></ellipse>
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

    this.state = {
      graph: sample,
      selected: {}
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
