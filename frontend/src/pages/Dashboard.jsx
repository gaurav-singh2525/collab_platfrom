import { useEffect, useState } from "react";

import { getCurrentUser } from "../services/authService";

import CodeEditor from "../components/CodeEditor";

import LanguageSelector from "../components/LanguageSelector";

import OutputPanel from "../components/OutputPanel";

import { executeCode } from "../services/executionService";

import socketService from "../services/socketService";

import { useParams } from "react-router-dom";

function Dashboard() {
  const [user, setUser] = useState(null);

  const [loading, setLoading] = useState(true);

  const [code, setCode] = useState(``);

  const [language, setLanguage] = useState("python");

  const [output, setOutput] = useState("");

  const [userCount, setUserCount] = useState(0);

  const [users, setUsers] = useState([]);

  const [activity, setActivity] = useState([]);

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
      const result = await executeCode(roomId, language, code);

      setOutput(result.stdout || result.stderr);
    } catch (err) {
      setOutput("Execution Failed");
    }
  };

  const handleLanguageChange = (newLanguage) => {
    setLanguage(newLanguage);

    socketService.send({
      type: "language_change",
      language: newLanguage,
    });
  };

  const { roomId } = useParams();

  useEffect(() => {
    const username = localStorage.getItem("email");
    const ws = socketService.connect(roomId, username);

    ws.onmessage = (event) => {
      const message = JSON.parse(event.data);

      if (message.type === "user_joined") {
        setActivity((prev) => {
          const updated = [...prev, `${message.username} joined room`];

          return updated.slice(-5);
        });
      }
      if (message.type === "user_left") {
        setActivity((prev) => {
          const updated = [...prev, `${message.username} joined room`];

          return updated.slice(-5);
        });
      }
      if (message.type === "code_change") {
        setCode(message.code);
      }
      if (message.type === "presence") {
        setUserCount(message.count);
        setUsers(message.users || []);
      }
      if (message.type === "language_change") {
        setLanguage(message.language);
      }
      if (message.type === "execution_output") {
        setOutput(message.output);
      }
    };

    return () => {
      ws.close();
    };
  }, []);

  useEffect(() => {
    fetchUser();
  }, []);

  if (loading) {
    return <h1>Loading...</h1>;
  }

  return (
    <div>
      <h1>Dashboard</h1>
      <p>Welcome: {user?.email}</p>
      <p>
        Connected Users:
        {userCount}
      </p>
      <div>
        <h3>Activity</h3>

        <ul>
          {activity.map((item, index) => (
            <li key={index}>{item}</li>
          ))}
        </ul>
      </div>
      <div>
        <h3>Participants</h3>

        <ul>
          {users.map((user) => (
            <li key={user}>{user}</li>
          ))}
        </ul>
      </div>
      <div>
        <h3>Room ID:</h3>
        <code>{roomId}</code>
      </div>

      <LanguageSelector
        language={language}
        onLanguageChange={handleLanguageChange}
      />

      <CodeEditor code={code} setCode={setCode} language={language} />

      <button onClick={handleRun}>Run Code</button>

      <OutputPanel output={output} />
    </div>
  );
}

export default Dashboard;
