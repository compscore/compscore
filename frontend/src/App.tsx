import { CssBaseline } from "@mui/material";
import { ThemeProvider } from "@mui/material/styles";
import { Suspense, useEffect, useState } from "react";
import { useCookies } from "react-cookie";
import { RouterProvider, createBrowserRouter } from "react-router-dom";

import Main from "./components/Main";
import Checks from "./pages/Checks";
import Index from "./pages/Index";
import Login from "./pages/Login";
import CheckScoreboard from "./pages/Scoreboard/Check";
import CheckRoundScoreboard from "./pages/Scoreboard/CheckRound";
import RoundScoreboard from "./pages/Scoreboard/Round";
import Scoreboard from "./pages/Scoreboard/Scoreboard";
import StatusScoreboard from "./pages/Scoreboard/Status";
import StatusRoundScoreboard from "./pages/Scoreboard/StatusRound";
import TeamScoreboard from "./pages/Scoreboard/Team";
import TeamRoundScoreboard from "./pages/Scoreboard/TeamRound";
import { useSystemTheme } from "./themes/Preference";
import ChangePassword from "./pages/ChangePassword";

const LazyComponent = ({
  element,
}: {
  element: React.ReactNode;
}): React.ReactElement => {
  return <Suspense fallback={<>Loading...</>}>{element}</Suspense>;
};

export default function App() {
  const [cookies, setCookie, removeCookie] = useCookies(["auth"]);
  const [mobile, setMobile] = useState<boolean>(window.innerWidth <= 768);

  const router = createBrowserRouter([
    {
      path: "/",
      element: (
        <LazyComponent
          element={
            <Main
              mobile={mobile}
              cookies={cookies}
              removeCookie={removeCookie}
            />
          }
        />
      ),
      errorElement: <Index />,
      children: [
        {
          index: true,
          element: <LazyComponent element={<Index />} />,
        },
        {
          path: "login",
          element: <LazyComponent element={<Login setCookie={setCookie} />} />,
        },
        {
          path: "password",
          element: <LazyComponent element={<ChangePassword />} />,
        },
        {
          path: "checks",
          element: <LazyComponent element={<Checks cookies={cookies} />} />,
        },
        {
          path: "scoreboard",
          element: <LazyComponent element={<Scoreboard />} />,
        },
        {
          path: "scoreboard/round/:round",
          element: <LazyComponent element={<RoundScoreboard />} />,
        },
        {
          path: "scoreboard/team/:team",
          element: <LazyComponent element={<TeamScoreboard />} />,
        },
        {
          path: "scoreboard/team/:team/:round",
          element: <LazyComponent element={<TeamRoundScoreboard />} />,
        },
        {
          path: "scoreboard/check/:check",
          element: <LazyComponent element={<CheckScoreboard />} />,
        },
        {
          path: "scoreboard/check/:check/:round",
          element: <LazyComponent element={<CheckRoundScoreboard />} />,
        },
        {
          path: "scoreboard/status/:team/:check",
          element: <LazyComponent element={<StatusScoreboard />} />,
        },
        {
          path: "scoreboard/status/:team/:check/:round",
          element: <LazyComponent element={<StatusRoundScoreboard />} />,
        },
      ],
    },
    {
      path: "/iframe",
      errorElement: <Index />,
      children: [
        {
          path: "scoreboard",
          element: <LazyComponent element={<Scoreboard />} />,
        },
        {
          path: "scoreboard/round/:round",
          element: <LazyComponent element={<RoundScoreboard />} />,
        },
        {
          path: "scoreboard/team/:team",
          element: <LazyComponent element={<TeamScoreboard />} />,
        },
        {
          path: "scoreboard/team/:team/:round",
          element: <LazyComponent element={<TeamRoundScoreboard />} />,
        },
        {
          path: "scoreboard/check/:check",
          element: <LazyComponent element={<CheckScoreboard />} />,
        },
        {
          path: "scoreboard/status/:team/:check",
          element: <LazyComponent element={<StatusScoreboard />} />,
        },
        {
          path: "scoreboard/status/:team/:check/:round",
          element: <LazyComponent element={<StatusRoundScoreboard />} />,
        },
      ],
    },
  ]);

  useEffect(() => {
    const updateWindowDimensions = () => {
      setMobile(window.innerWidth <= 600);
    };

    window.addEventListener("resize", updateWindowDimensions);

    return () => {
      window.removeEventListener("resize", updateWindowDimensions);
    };
  }, []);

  return (
    <ThemeProvider theme={useSystemTheme()}>
      <CssBaseline />
      <RouterProvider router={router} />
    </ThemeProvider>
  );
}
