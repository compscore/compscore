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
import { Scoreboard } from "../../models/Scoreboard/Scoreboard";
import { Round } from "../../models/ent";
import DoubleArrowLeftIcon from "@mui/icons-material/KeyboardDoubleArrowLeft";
import DoubleArrowRightIcon from "@mui/icons-material/KeyboardDoubleArrowRight";
import ArrowLeftIcon from "@mui/icons-material/KeyboardArrowLeft";
import ArrowRightIcon from "@mui/icons-material/KeyboardArrowRight";
import { api_url } from "../../config";

type props = {
  round: string;
};

export default function RoundScoreboardComponent({ round }: props) {
  const [data, setData] = useState<Scoreboard>();
  const [latestRound, setLatestRound] = useState<Round>();

  useEffect(() => {
    const fetchData = async () => {
      fetch(`${api_url}/api/round/latest`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then(async (res) => {
          if (res.status === 200) {
            const response = (await res.json()) as Round;

            setLatestRound(response);

            if (0 >= parseInt(round) || parseInt(round) >= response.number) {
              window.location.href = "/scoreboard";
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
      fetch(`${api_url}/api/scoreboard/round/${round}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then(async (res) => {
          if (res.status === 200) {
            setData((await res.json()) as Scoreboard);
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

    const pollingInterval = setInterval(fetchData, 15000);

    return () => clearInterval(pollingInterval);
  }, []);

  const [highlightedTeam, setHighlightedTeam] = useState<number | null>(null);
  const [highlightedCheck, setHighlightedCheck] = useState<string | null>(null);

  const getBackgroundColor = (status: number, team: number, check: string) => {
    if (highlightedCheck === null && highlightedTeam === null) {
      if (status === 0) {
        return "#f44336";
      } else if (status === 1) {
        return "#4caf50";
      }
      return "#999891";
    }
    if (highlightedCheck === check || highlightedTeam === team) {
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
        Scoreboard
      </Typography>
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
        }}
      >
        {parseInt(round) >= 10 ? (
          <DoubleArrowLeftIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href =
                "/scoreboard/round/" + (parseInt(round) - 10);
            }}
          />
        ) : (
          <ArrowLeftIcon
            sx={{
              visibility: "hidden",
            }}
          />
        )}
        {parseInt(round) >= 1 ? (
          <ArrowLeftIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href =
                "/scoreboard/round/" + (parseInt(round) - 1);
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
            window.location.href = "/scoreboard";
          }}
        >
          Round {data?.round}
        </Typography>
        {latestRound && parseInt(round) <= latestRound?.number ? (
          <ArrowRightIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href =
                "/scoreboard/round/" + (parseInt(round) + 1);
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
              window.location.href =
                "/scoreboard/round/" + (parseInt(round) + 10);
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
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell size='small'>
                <Typography variant='subtitle2'>Team</Typography>
              </TableCell>
              {data?.checks[0].status.map((_, team) => (
                <TableCell
                  key={team + 1}
                  align='center'
                  size='small'
                  onMouseEnter={() => {
                    setHighlightedTeam(team + 1);
                  }}
                  onMouseLeave={() => {
                    setHighlightedTeam(null);
                  }}
                  sx={{
                    backgroundColor:
                      highlightedTeam === team + 1 ? "#343434" : "transparent",
                  }}
                  onClick={() => {
                    window.location.href = `/scoreboard/team/${
                      team + 1
                    }/${round}`;
                  }}
                >
                  <Typography variant='subtitle2'>{team + 1}</Typography>
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {data?.checks.map((check) => (
              <TableRow>
                <TableCell
                  key={check.name}
                  size='small'
                  onMouseEnter={() => {
                    setHighlightedCheck(check.name);
                  }}
                  onMouseLeave={() => {
                    setHighlightedCheck(null);
                  }}
                  sx={{
                    backgroundColor:
                      highlightedCheck === check.name
                        ? "#343434"
                        : "transparent",
                  }}
                  onClick={() => {
                    window.location.href = `/scoreboard/check/${check.name}/${round}`;
                  }}
                >
                  {check.name}
                </TableCell>
                {check.status.map((status, team) => (
                  <TableCell
                    key={team + "-" + check.name}
                    size='small'
                    sx={{
                      backgroundColor: getBackgroundColor(
                        status,
                        team + 1,
                        check.name
                      ),
                    }}
                    onMouseEnter={() => {
                      setHighlightedTeam(team + 1);
                      setHighlightedCheck(check.name);
                    }}
                    onMouseLeave={() => {
                      setHighlightedTeam(null);
                      setHighlightedCheck(null);
                    }}
                    onClick={() => {
                      window.location.href = `/scoreboard/status/${team + 1}/${
                        check.name
                      }/${round}`;
                    }}
                  ></TableCell>
                ))}
              </TableRow>
            ))}
            <TableRow>
              <TableCell size='small'>
                <Typography variant='subtitle2'>Score</Typography>
              </TableCell>
              {data?.scores.map((score, team) => (
                <TableCell
                  key={"score" + team + 1}
                  size='small'
                  align='center'
                  onMouseEnter={() => {
                    setHighlightedTeam(team + 1);
                  }}
                  onMouseLeave={() => {
                    setHighlightedTeam(null);
                  }}
                  sx={{
                    backgroundColor:
                      highlightedTeam === team + 1 ? "#343434" : "transparent",
                  }}
                >
                  <Typography variant='subtitle2'>{score}</Typography>
                </TableCell>
              ))}
            </TableRow>
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}
