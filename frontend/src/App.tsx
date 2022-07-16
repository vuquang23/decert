import HomePage from "./pages/HomePage";
import { Routes, Route } from "react-router-dom";
import IssuePage from "./pages/IssuePage";

const App = () => (
  <Routes>
    <Route path="/" element={<HomePage />} />
    <Route path="/issue/:page" element={<IssuePage />} />
  </Routes>
);
export default App;
