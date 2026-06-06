import { useEffect, useState } from "react";

import { getCurrentUser } from "../services/authService";

function Dashboard() {
  const [user, setUser] = useState(null);

  const [loading, setLoading] = useState(true);

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

  useEffect(() => {
    fetchUser();
  }, []);

  if (loading) {
    return <h1>Loading...</h1>;
  }

  return (
    <div>
      <h1>Dashboard</h1>

      <p>
        Welcome:
        {user?.email}
      </p>
    </div>
  );
}

export default Dashboard;
