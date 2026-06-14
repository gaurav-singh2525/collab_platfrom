import { useState } from "react";
import { useNavigate } from "react-router-dom";

import { signup } from "../services/authService";

function Signup() {
  const [email, setEmail] = useState("");

  const [password, setPassword] = useState("");

  const [loading, setLoading] = useState(false);

  const [error, setError] = useState("");

  const navigate = useNavigate();

  const handleSignup = async () => {
    setError("");

    if (!email || !password) {
      setError("Please enter email and password");

      return;
    }

    try {
      setLoading(true);

      await signup(email, password);

      navigate("/login");
    } catch (err) {
      setError(err.response?.data?.error || "Signup failed");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-slate-950 flex items-center justify-center px-4">
      <div className="w-full max-w-md rounded-2xl border border-slate-800 bg-slate-900 p-8 shadow-xl">
        <h1 className="text-3xl font-bold text-white text-center mb-2">
          Create Account
        </h1>

        <p className="text-slate-400 text-center mb-8">
          Start collaborating in real time
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
            onClick={handleSignup}
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
            {loading ? "Creating Account..." : "Sign Up"}
          </button>

          {error && <p className="text-red-400 text-sm text-center">{error}</p>}
        </div>

        <p className="mt-6 text-center text-sm text-slate-400">
          Already have an account?
          <span
            onClick={() => navigate("/login")}
            className="
              ml-2
              cursor-pointer
              text-green-400
              hover:text-green-300
            "
          >
            Login
          </span>
        </p>
      </div>
    </div>
  );
}

export default Signup;
