function LanguageSelector({ language, onLanguageChange }) {
  return (
    <select
      value={language}
      onChange={(e) => onLanguageChange(e.target.value)}
      className="
    rounded-lg
    bg-slate-800
    border
    border-slate-700
    px-3
    py-2
    text-white
    outline-none
  "
    >
      <option value="python">Python</option>

      <option value="cpp">C++</option>

      <option value="javascript">JavaScript</option>
    </select>
  );
}

export default LanguageSelector;
