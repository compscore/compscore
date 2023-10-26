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
import { CheckScoreboard } from "../../models/Scoreboard/Check";
import { Round } from "../../models/ent";
import ArrowLeftIcon from "@mui/icons-material/KeyboardArrowLeft";
import ArrowRightIcon from "@mui/icons-material/KeyboardArrowRight";
import DoubleArrowLeftIcon from "@mui/icons-material/KeyboardDoubleArrowLeft";
import DoubleArrowRightIcon from "@mui/icons-material/KeyboardDoubleArrowRight";
import { api_url } from "../../config";

type props = {
  check: string;
  round: string;
};

export default function CheckRoundScoreboardComponent({ check, round }: props) {
  const [data, setData] = useState<CheckScoreboard>();
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
              window.location.href = `/scoreboard/check/${check}`;
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
      fetch(`${api_url}/api/scoreboard/check/${check}/${round}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then(async (res) => {
          if (res.status === 200) {
            setData((await res.json()) as CheckScoreboard);
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

    return () => clearInterval(pollingInterval);
  }, []);

  const [highlightedRound, setHighlightedRound] = useState<number | null>(null);
  const [highlightedTeam, setHighlightedTeam] = useState<string | null>(null);

  const getBackgroundColor = (status: number, round: number, check: string) => {
    if (highlightedTeam === null && highlightedRound === null) {
      if (status === 0) {
        return "#f44336";
      } else if (status === 1) {
        return "#4caf50";
      }
      return "#999891";
    }
    if (highlightedTeam === check || highlightedRound === round) {
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
        {check}
      </Typography>
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
        }}
      >
        {data && data.round > 10 ? (
          <DoubleArrowLeftIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href = `/scoreboard/check/${check}/${
                data.round - 10
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
        {data && data.round > 1 ? (
          <ArrowLeftIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href = `/scoreboard/check/${check}/${
                data.round - 1
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
            window.location.href = `/scoreboard/check/${check}`;
          }}
        >
          Round {round}
        </Typography>
        {latestRound && data && data.round < latestRound?.number ? (
          <ArrowRightIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href = `/scoreboard/check/${check}/${
                data.round + 1
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
        {latestRound && data && data.round + 10 <= latestRound?.number ? (
          <DoubleArrowRightIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href = `/scoreboard/check/${check}/${
                data.round + 10
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
              {data?.teams[0].status.map((_, index) => (
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
            {data?.teams.map((team, index) => (
              <TableRow key={index}>
                <TableCell
                  size='small'
                  sx={{
                    backgroundColor:
                      highlightedTeam === team.name ? "#343434" : "transparent",
                  }}
                  onMouseEnter={() => {
                    setHighlightedTeam(team.name);
                  }}
                  onMouseLeave={() => {
                    setHighlightedTeam(null);
                  }}
                  onClick={() => {
                    window.location.href = `/scoreboard/team/${
                      index + 1
                    }/${round}`;
                  }}
                >
                  {team.name}
                </TableCell>
                {team.status.map((status, s_index) => (
                  <TableCell
                    key={s_index}
                    size='small'
                    sx={{
                      backgroundColor: getBackgroundColor(
                        status,
                        s_index,
                        team.name
                      ),
                    }}
                    onMouseEnter={() => {
                      setHighlightedRound(s_index);
                      setHighlightedTeam(team.name);
                    }}
                    onMouseLeave={() => {
                      setHighlightedRound(null);
                      setHighlightedTeam(null);
                    }}
                    onClick={() => {
                      window.location.href = `/scoreboard/status/${
                        index + 1
                      }/${check}/${data?.round - s_index}`;
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
