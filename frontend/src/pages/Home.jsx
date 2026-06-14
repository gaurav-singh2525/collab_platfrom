import { useNavigate } from "react-router-dom";
import { useState } from "react";

function Home() {
  const navigate = useNavigate();
  const [roomLink, setRoomLink] = useState("");
  const createRoom = () => {
    const roomId = crypto.randomUUID();

    navigate(`/editor/${roomId}`);
  };

  const joinRoom = () => {
    const value = roomLink.trim();

    if (!value) return;

    try {
      if (value.startsWith("http")) {
        const url = new URL(value);

        const parts = url.pathname.split("/");

        const roomId = parts[parts.length - 1];

        navigate(`/editor/${roomId}`);
      } else {
        navigate(`/editor/${value}`);
      }
    } catch {
      alert("Invalid room link");
    }
  };

  return (
    <div className="min-h-screen bg-slate-950 text-white flex items-center justify-center px-4">
      <div className="max-w-3xl text-center">
        <h1 className="text-5xl font-bold mb-6">Collaborative Code Platform</h1>

        <p className="text-slate-400 text-lg mb-8">
          Real-time collaborative coding with multi-language execution, room
          persistence, live presence tracking, and Docker sandboxing.
        </p>

        <button
          onClick={createRoom}
          className="
            rounded-xl
            bg-green-600
            px-8
            py-4
            text-lg
            font-medium
            transition
            hover:bg-green-700
          "
        >
          Create New Room
        </button>
        <div className="mt-6">
          <p className="mb-3 text-slate-400">Or join an existing room</p>

          <div className="flex gap-2 justify-center">
            <input
              type="text"
              placeholder="Paste room link"
              value={roomLink}
              onChange={(e) => setRoomLink(e.target.value)}
              onKeyDown={(e) => {
                if (e.key === "Enter") {
                  joinRoom();
                }
              }}
              className="
        w-96
        rounded-lg
        border
        border-slate-700
        bg-slate-800
        px-4
        py-3
        text-white
      "
            />

            <button
              onClick={joinRoom}
              className="
        rounded-lg
        bg-blue-600
        px-5
        py-3
        hover:bg-blue-700
      "
            >
              Join Room
            </button>
          </div>
        </div>
        <div className="mt-12 grid grid-cols-1 md:grid-cols-3 gap-4">
          <div className="rounded-xl border border-slate-800 bg-slate-900 p-4">
            <h3 className="font-semibold mb-2">Real-Time Collaboration</h3>

            <p className="text-sm text-slate-400">
              Code changes sync instantly across all connected users.
            </p>
          </div>

          <div className="rounded-xl border border-slate-800 bg-slate-900 p-4">
            <h3 className="font-semibold mb-2">Multi-Language Support</h3>

            <p className="text-sm text-slate-400">
              Run Python, JavaScript, and C++ code securely.
            </p>
          </div>

          <div className="rounded-xl border border-slate-800 bg-slate-900 p-4">
            <h3 className="font-semibold mb-2">Persistent Rooms</h3>

            <p className="text-sm text-slate-400">
              Rejoin rooms anytime and continue where you left off.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Home;
