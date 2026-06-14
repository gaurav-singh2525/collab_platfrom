import { useEffect, useState } from "react";

import { getCurrentUser } from "../services/authService";

import CodeEditor from "../components/CodeEditor";

import LanguageSelector from "../components/LanguageSelector";

import { useAuth } from "../context/AuthContext";

import { useNavigate } from "react-router-dom";

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

  const [copied, setCopied] = useState(false);

  const [input, setInput] = useState("");

  const auth = useAuth();
  const navigate = useNavigate();

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
      const result = await executeCode(roomId, language, code, input);

      setOutput(result.stdout || result.stderr);
    } catch (err) {
      setOutput(err.response?.data?.error || "Execution Failed");
    }
  };

  const handleCopy = async () => {
    await navigator.clipboard.writeText(window.location.href);

    setCopied(true);

    setTimeout(() => {
      setCopied(false);
    }, 2000);
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
          const updated = [...prev, `${message.username} left room`];

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
    <div className="min-h-screen bg-slate-950 text-white">
      {/* Header */}
      <div className="border-b border-slate-800 bg-slate-900 px-6 py-4">
        <div className="flex flex-col gap-3 lg:flex-row lg:items-center lg:justify-between">
          <div>
            <h1 className="text-2xl font-bold">Collaborative Code Platform</h1>

            <p className="mt-1 text-sm text-slate-400">{user?.email}</p>
          </div>
          <div className="text-right text-sm">
            <button
              onClick={() => {
                auth.logout();
                navigate("/login");
              }}
              className="
    mb-2
    rounded-lg
    bg-red-600
    px-3
    py-2
    text-white
    hover:bg-red-700
  "
            >
              Logout
            </button>
          </div>
          <div className="flex items-center gap-2">
            <span className="text-slate-300">Room:</span>

            <span className="font-medium text-white">
              {roomId?.slice(0, 8) || "Loading..."}
            </span>

            <button
              onClick={handleCopy}
              className="
              rounded
              bg-slate-700
              px-2
              py-1
              text-xs
              hover:bg-slate-600
            "
            >
              {copied ? "Copied!" : "Copy Link"}
            </button>
          </div>
        </div>
      </div>

      {/* Main Layout */}
      <div className="flex flex-col lg:flex-row gap-4 p-4">
        {/* Sidebar */}
        <div
          className="
          w-full
          lg:w-64
          flex-shrink-0
          space-y-4
          lg:max-h-[calc(100vh-120px)]
          lg:overflow-y-auto
        "
        >
          {/* Participants */}
          <div className="rounded-xl border border-slate-800 bg-slate-900 p-4">
            <h3 className="mb-3 text-lg font-semibold">
              Participants ({userCount})
            </h3>

            <ul className="space-y-2">
              {users.map((user) => (
                <li
                  key={user}
                  className="
                  flex
                  items-center
                  gap-2
                  break-all
                  text-slate-200
                "
                >
                  <span className="text-green-400">●</span>
                  {user}
                </li>
              ))}
            </ul>
          </div>

          {/* Activity */}
          <div className="rounded-xl border border-slate-800 bg-slate-900 p-4">
            <h3 className="mb-3 text-lg font-semibold">Activity</h3>

            <div className="max-h-64 overflow-y-auto">
              <ul className="space-y-2">
                {activity.map((item, index) => (
                  <li
                    key={index}
                    className="
                    break-words
                    text-sm
                    text-slate-300
                  "
                  >
                    {item}
                  </li>
                ))}
              </ul>
            </div>
          </div>
        </div>

        {/* Main Section */}
        <div className="flex-1 min-w-0 space-y-4">
          {/* Toolbar */}
          <div
            className="
            flex
            flex-col
            gap-3
            sm:flex-row
            sm:items-center
            sm:justify-between
            rounded-xl
            border
            border-slate-800
            bg-slate-900
            p-4
          "
          >
            <LanguageSelector
              language={language}
              onLanguageChange={handleLanguageChange}
            />

            <button
              onClick={handleRun}
              className="
              rounded-lg
              bg-green-600
              px-5
              py-2
              font-medium
              transition
              hover:bg-green-700
            "
            >
              Run Code
            </button>
          </div>

          {/* Editor */}
          <div
            className="
            h-[500px]
            overflow-hidden
            rounded-xl
            border
            border-slate-800
          "
          >
            <CodeEditor code={code} setCode={setCode} language={language} />
          </div>

          {/* input */}
          
          <div className="rounded-xl border border-slate-800 bg-slate-900 p-4">
            <h3 className="mb-3 text-lg font-semibold">Input</h3>

            <textarea
              value={input}
              onChange={(e) => setInput(e.target.value)}
              placeholder="Enter program input here..."
              className="
      w-full
      h-32
      resize-none
      rounded-lg
      bg-slate-950
      p-3
      text-white
      border
      border-slate-700
      outline-none
    "
            />
          </div>

          {/* Output */}
          <div className="rounded-xl border border-slate-800 bg-slate-900 p-4">
            <h3 className="mb-3 text-lg font-semibold">Output</h3>

            <div
              className="
      h-48
      overflow-y-auto
      overflow-x-auto
      rounded-lg
      bg-slate-950
      p-3
      text-slate-200
      whitespace-pre-wrap
      break-words
    "
            >
              <OutputPanel output={output} />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
