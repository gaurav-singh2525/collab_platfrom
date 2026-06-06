import {
  Navigate,
} from "react-router-dom";

import {
  useAuth,
} from "../context/AuthContext";

function ProtectedRoute({
  children,
}) {

  const auth =
    useAuth();

  if (
    !auth.isAuthenticated
  ) {
    return (
      <Navigate
        to="/login"
      />
    );
  }

  return children;
}

export default ProtectedRoute;