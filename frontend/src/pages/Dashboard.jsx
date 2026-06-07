import { useEffect, useState } from "react";

import { getCurrentUser } from "../services/authService";

import CodeEditor from "../components/CodeEditor";

import LanguageSelector from "../components/LanguageSelector";

import OutputPanel from "../components/OutputPanel";

import { executeCode } from "../services/executionService";

function Dashboard() {
  const [user, setUser] = useState(null);

  const [loading, setLoading] = useState(true);

  const [code, setCode] = useState(`print("HelloWolrd")`);

  const [language, setLanguage] = useState("python");

  const [output, setOutput] = useState("");

  const fetchUser = async () => {
    try {
      const data = await getCurrentUser();

      setUser(data);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleRun = async () => {
    try {
      const result = await executeCode(language, code);

      setOutput(result.stdout || result.stderr);
    } catch (err) {
      setOutput("Execution Failed");
    }
  };

  useEffect(() => {
    fetchUser();
  }, []);

  if (loading) {
    return <h1>Loading...</h1>;
  }

  return (
    <div>
      <h1>Dashboard</h1>
      <p>Welcome: {user?.Email}</p>
      <LanguageSelector language={language} setLanguage={setLanguage} />

      <CodeEditor code={code} setCode={setCode} language={language} />

      <button onClick={handleRun} >Run Code</button>

      <OutputPanel output={output} />
    </div>
  );
}

export default Dashboard;
