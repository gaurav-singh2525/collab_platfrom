import { useState } from "react";
import { useNavigate } from "react-router-dom";

import { login } from "../services/authService";
import { useAuth } from "../context/AuthContext";

function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const auth = useAuth();
  const navigate = useNavigate();

  const handleLogin = async () => {
    setError("");
    if (!email || !password) {
      setError("Please enter email and password");
      return;
    }

    try {
      setLoading(true);

      const data = await login(email, password);

      auth.login(data.token, data.email);

      navigate("/");
    } catch (err) {
      setError(err.response?.data?.error || "Login failed");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-slate-950 flex items-center justify-center px-4">
      <div className="w-full max-w-md rounded-2xl border border-slate-800 bg-slate-900 p-8 shadow-xl">
        <h1 className="text-3xl font-bold text-white text-center mb-2">
          Welcome Back
        </h1>

        <p className="text-slate-400 text-center mb-8">
          Sign in to continue coding collaboratively
        </p>

        <div className="space-y-4">
          <input
            type="email"
            placeholder="Enter email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="
            w-full
            rounded-lg
            border
            border-slate-700
            bg-slate-800
            px-4
            py-3
            text-white
            outline-none
            focus:border-green-500
          "
          />

          <input
            type="password"
            placeholder="Enter password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="
            w-full
            rounded-lg
            border
            border-slate-700
            bg-slate-800
            px-4
            py-3
            text-white
            outline-none
            focus:border-green-500
          "
          />

          <button
            onClick={handleLogin}
            disabled={loading}
            className="
            w-full
            rounded-lg
            bg-green-600
            py-3
            font-medium
            text-white
            transition
            hover:bg-green-700
            disabled:opacity-50
          "
          >
            {loading ? "Logging in..." : "Login"}
          </button>

          {error && <p className="text-red-400 text-sm text-center">{error}</p>}
        </div>

        <p className="mt-6 text-center text-sm text-slate-400">
          Don't have an account?
          <span
            onClick={() => navigate("/signup")}
            className="
            ml-2
            cursor-pointer
            text-green-400
            hover:text-green-300
          "
          >
            Sign Up
          </span>
        </p>
      </div>
    </div>
  );
}

export default Login;
