import React, { Component } from 'react'
import MonacoEditor from 'react-monaco-editor'

class JSONEditor extends Component {

  constructor(props) {
    super()

    this.state = {
      code: props.code,
    }

    this.onChange = this.onChange.bind(this)
    this.editorDidMount = this.editorDidMount.bind(this)
  }

  editorDidMount(editor, monaco) {
    console.log('editorDidMount', editor);
    editor.focus();
  }

  onChange(newValue, e) {
    console.log('onChange', newValue, e);
  }

  render() {
    return (
      <MonacoEditor
        language="json"
        theme="vs-dark"
        value={JSON.stringify(this.state.code, null, 2)}
        // options={options}
        onChange={this.onChange}
        editorDidMount={this.editorDidMount}
      />
    )
  }
}

export default JSONEditor
