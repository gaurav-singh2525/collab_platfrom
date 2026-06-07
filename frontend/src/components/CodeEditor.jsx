import Editor from "@monaco-editor/react";

function CodeEditor({ code, setCode, language }) {
  return (
    <Editor
      height="500px"
      language={language}
      value={code}
      onChange={(value) => setCode(value)}
      theme="vs-dark"
    />
  );
}

export default CodeEditor;
