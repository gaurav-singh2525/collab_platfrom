function LanguageSelector({ language, onLanguageChange }) {
  return (
    <select value={language} onChange={(e) => onLanguageChange(e.target.value)}>
      <option value="python">Python</option>

      <option value="cpp">C++</option>

      <option value="javascript">JavaScript</option>
    </select>
  );
}

export default LanguageSelector;
