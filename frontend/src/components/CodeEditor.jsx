import Editor from "@monaco-editor/react";
import socketService from "../services/socketService";

function CodeEditor({ code, setCode, language }) {
  return (
    <Editor
      height="500px"
      language={language}
      value={code}
      onChange={(value) => {
        setCode(value);

        socketService.send({
          type: "code_change",
          code: value,
        });
      }}
      theme="vs-dark"
    />
  );
}

export default CodeEditor;
