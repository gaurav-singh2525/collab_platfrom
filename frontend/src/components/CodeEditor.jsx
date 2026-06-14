import Editor from "@monaco-editor/react";
import { useRef } from "react";
import socketService from "../services/socketService";

function CodeEditor({ code, setCode, language }) {
  const timeoutRef = useRef(null);

  const handleChange = (value = "") => {
    setCode(value);

    clearTimeout(timeoutRef.current);

    timeoutRef.current = setTimeout(() => {
      socketService.send({
        type: "code_change",
        code: value,
      });
    }, 50);
  };

  return (
    <Editor
      height="75vh"
      language={language}
      value={code}
      onChange={handleChange}
      theme="vs-dark"
      options={{
        minimap: { enabled: false },

        fontSize: 20,
        fontLigatures: true,

        wordWrap: "on",

        automaticLayout:true,

        tabSize: 2,

        scrollBeyondLastLine: false,

        renderWhitespace: "selection",

        renderLineHighlight: "all",

        cursorBlinking: "smooth",

        cursorSmoothCaretAnimation: "on",

        cursorStyle: "line",

        smoothScrolling: true,

        matchBrackets: "always",

        autoClosingBrackets: "always",

        autoClosingQuotes: "always",

        autoIndent: "full",

        quickSuggestions: true,

        suggestOnTriggerCharacters: true,

        acceptSuggestionOnEnter: "on",

        guides: {
          bracketPairs: true,
        },

        scrollbar: {
          verticalScrollbarSize: 8,
          horizontalScrollbarSize: 8,
        },

        padding: {
          top: 12,
        },
      }}
    />
  );
}

export default CodeEditor;
