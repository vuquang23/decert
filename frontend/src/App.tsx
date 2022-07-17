import HomePage from "pages/HomePage";
import IssuePage from "pages/IssuePage";
import { Route, Routes } from "react-router-dom";

const App = () => (
  <Routes>
    <Route path="/" element={<HomePage />} />
    <Route path="/issue/:page" element={<IssuePage />} />
  </Routes>
);
export default App;
