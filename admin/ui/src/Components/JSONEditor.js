import React, { Component } from 'react'
import MonacoEditor from 'react-monaco-editor'

class JSONEditor extends Component {
  static getDerivedStateFromProps(props, state) { // eslint-disable-line
    return {
      code: props.code,
    }
  }

  constructor(props) { // eslint-disable-line
    super()

    this.state = {
      code: '',
    }

    this.onChange = this.onChange.bind(this)
    this.editorDidMount = this.editorDidMount.bind(this)
  }

  onChange(newValue, e) { // eslint-disable-line
    try {
      const response = JSON.parse(newValue)
      if (response._id !== undefined) {
        this.props.onChange(response) // eslint-disable-line
      }
    } catch (e) { // eslint-disable-line
      // ignore update if it's invalid JSON
    }
  }

  editorDidMount(editor, monaco) { // eslint-disable-line
    // console.warn('editorDidMount', editor);
    editor.focus()
  }

  render() {
    return (
      <MonacoEditor
        language="json"
        theme="vs-dark"
        value={JSON.stringify(this.state.code, null, 2)} // eslint-disable-line
        // options={options}
        onChange={this.onChange}
        editorDidMount={this.editorDidMount}
      />
    )
  }
}

export default JSONEditor
