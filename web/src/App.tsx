import './App.css';

// eslint-disable-next-line react-refresh/only-export-components
export const API_URL =
  import.meta.env.VITE_API_URL || 'http://localhost:8080/v1';

function App() {
  return (
    <>
      <div>Home screen</div>
    </>
  );
}

export default App;
