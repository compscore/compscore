import {
  Box,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";
import { enqueueSnackbar } from "notistack";
import { useEffect, useState } from "react";
import { TeamScoreboard } from "../../models/Scoreboard/Team";
import { Round } from "../../models/ent";
import ArrowLeftIcon from "@mui/icons-material/KeyboardArrowLeft";
import ArrowRightIcon from "@mui/icons-material/KeyboardArrowRight";
import DoubleArrowLeftIcon from "@mui/icons-material/KeyboardDoubleArrowLeft";
import DoubleArrowRightIcon from "@mui/icons-material/KeyboardDoubleArrowRight";

type props = {
  team: string;
  round: string;
};

export default function TeamRoundScoreboardComponent({ team, round }: props) {
  const [data, setData] = useState<TeamScoreboard>();
  const [latestRound, setLatestRound] = useState<Round>();

  useEffect(() => {
    const fetchData = async () => {
      fetch(`http://localhost:8080/api/round/latest`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then(async (res) => {
          if (res.status === 200) {
            let response = (await res.json()) as Round;

            setLatestRound(response);

            if (0 >= parseInt(round) || parseInt(round) >= response.number) {
              window.location.href = `/scoreboard/team/${team}`;
            }
          } else {
            enqueueSnackbar("Encountered an error", { variant: "error" });
          }
        })
        .catch((err) => {
          enqueueSnackbar("Encountered an error: " + err, { variant: "error" });
          console.log(err);
        });
    };

    fetchData();
  }, []);

  useEffect(() => {
    const fetchData = async () => {
      fetch(`http://localhost:8080/api/scoreboard/team/${team}/${round}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then(async (res) => {
          let response = (await res.json()) as TeamScoreboard;
          if (res.status === 200) {
            setData(response);
          } else {
            enqueueSnackbar("Encountered an error", { variant: "error" });
          }
        })
        .catch((err) => {
          enqueueSnackbar("Encountered an error: " + err, { variant: "error" });
          console.log(err);
        });
    };

    fetchData();

    const pollingInterval = setInterval(fetchData, 5000);

    console.log(data);
    return () => clearInterval(pollingInterval);
  }, []);

  const [highlightedRound, setHighlightedRound] = useState<number | null>(null);
  const [highlightedCheck, setHighlightedCheck] = useState<string | null>(null);

  const getBackgroundColor = (status: number, round: number, check: string) => {
    if (highlightedCheck === null && highlightedRound === null) {
      if (status === 0) {
        return "#f44336";
      } else if (status === 1) {
        return "#4caf50";
      }
      return "#999891";
    }
    if (highlightedCheck === check || highlightedRound === round) {
      if (status === 0) {
        return "#f44336";
      } else if (status === 1) {
        return "#4caf50";
      }
      return "#999891";
    }
    if (status === 0) {
      return "#f26359";
    } else if (status === 1) {
      return "#75ad77";
    }
    return "#cecdc6";
  };

  return (
    <>
      <Typography
        component='h1'
        variant='h3'
        fontWeight={700}
        sx={{
          marginTop: 5,
        }}
        onClick={() => {
          window.location.href = "/scoreboard";
        }}
      >
        Team {team}
      </Typography>
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
        }}
      >
        {parseInt(round) > 10 ? (
          <DoubleArrowLeftIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href = `/scoreboard/team/${team}/${
                parseInt(round) - 10
              }`;
            }}
          />
        ) : (
          <ArrowLeftIcon
            sx={{
              visibility: "hidden",
            }}
          />
        )}
        {parseInt(round) > 1 ? (
          <ArrowLeftIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href = `/scoreboard/team/${team}/${
                parseInt(round) - 1
              }`;
            }}
          />
        ) : (
          <ArrowLeftIcon
            sx={{
              visibility: "hidden",
            }}
          />
        )}
        <Typography
          component='h1'
          variant='h5'
          onClick={() => {
            window.location.href = `/scoreboard/team/${team}`;
          }}
        >
          Round {round}
        </Typography>
        {latestRound && parseInt(round) < latestRound?.number ? (
          <ArrowRightIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href = `/scoreboard/team/${team}/${
                parseInt(round) + 1
              }`;
            }}
          />
        ) : (
          <ArrowLeftIcon
            sx={{
              visibility: "hidden",
            }}
          />
        )}
        {latestRound && parseInt(round) + 10 <= latestRound?.number ? (
          <DoubleArrowRightIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href = `/scoreboard/team/${team}/${
                parseInt(round) + 10
              }`;
            }}
          />
        ) : (
          <ArrowLeftIcon
            sx={{
              visibility: "hidden",
            }}
          />
        )}
      </Box>
      <Box m={2}></Box>
      <TableContainer component={Paper} sx={{ width: "80%" }}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell size='small'>Round</TableCell>
              {data?.checks[0].status.map((_, index) => (
                <TableCell
                  size='small'
                  key={"round-" + (data?.round - index)}
                  sx={{
                    backgroundColor:
                      highlightedRound === index ? "#343434" : "transparent",
                  }}
                  onMouseEnter={() => {
                    setHighlightedRound(index);
                  }}
                  onMouseLeave={() => {
                    setHighlightedRound(null);
                  }}
                  onClick={() => {
                    window.location.href =
                      "/scoreboard/round/" + (data?.round - index);
                  }}
                >
                  {data?.round - index}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {data?.checks.map((check, index) => (
              <TableRow key={index}>
                <TableCell
                  size='small'
                  sx={{
                    backgroundColor:
                      highlightedCheck === check.name
                        ? "#343434"
                        : "transparent",
                  }}
                  onMouseEnter={() => {
                    setHighlightedCheck(check.name);
                  }}
                  onMouseLeave={() => {
                    setHighlightedCheck(null);
                  }}
                  onClick={() => {
                    window.location.href = `/scoreboard/check/${check.name}/${round}`;
                  }}
                >
                  {check.name}
                </TableCell>
                {check.status.map((status, index) => (
                  <TableCell
                    key={index}
                    size='small'
                    sx={{
                      backgroundColor: getBackgroundColor(
                        status,
                        index,
                        check.name
                      ),
                    }}
                    onMouseEnter={() => {
                      setHighlightedRound(index);
                      setHighlightedCheck(check.name);
                    }}
                    onMouseLeave={() => {
                      setHighlightedRound(null);
                      setHighlightedCheck(null);
                    }}
                    onClick={() => {
                      window.location.href = `/scoreboard/status/${team}/${
                        check.name
                      }/${data?.round - index}`;
                    }}
                  ></TableCell>
                ))}
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}
