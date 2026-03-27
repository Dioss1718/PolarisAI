import React, { useState } from "react";
import SplashIntro from "./components/SplashIntro";
import ConsolePage from "./pages/ConsolePage";

export default function App() {
  const [entered, setEntered] = useState(false);

  return entered ? (
    <ConsolePage />
  ) : (
    <SplashIntro onEnter={() => setEntered(true)} />
  );
}