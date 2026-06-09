import { useNavigate } from "react-router-dom";

function Home() {
  const navigate = useNavigate();

  const createRoom = () => {
    const roomId = crypto.randomUUID();

    navigate(`/editor/${roomId}`);
  };

  return (
    <div>
      <h1>Collaborative Code Platform</h1>

      <button onClick={createRoom}>Create Room</button>
    </div>
  );
}

export default Home;
