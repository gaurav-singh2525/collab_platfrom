function OutputPanel({ output }) {
  return (
    <div>
      <pre
        className="
  h-56
  overflow-y-auto
  overflow-x-auto
  whitespace-pre-wrap
  break-words
"
      >
        {output}
      </pre>
    </div>
  );
}

export default OutputPanel;
