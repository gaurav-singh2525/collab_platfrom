    function LanguageSelector({ language, setLanguage }) {
    return (
        <select value={language} onChange={(e) => setLanguage(e.target.value)}>
        <option value="python">Python</option>

        <option value="cpp">C++</option>

        <option value="javascript">JavaScript</option>
        </select>
    );
    }

    export default LanguageSelector;
