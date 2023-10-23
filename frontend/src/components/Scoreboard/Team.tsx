import ArrowLeftIcon from "@mui/icons-material/KeyboardArrowLeft";
import DoubleArrowLeftIcon from "@mui/icons-material/KeyboardDoubleArrowLeft";
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
import { api_url } from "../../config";

type props = {
  team: string;
};

export default function TeamScoreboardComponent({ team }: props) {
  const [data, setData] = useState<TeamScoreboard>();

  useEffect(() => {
    const fetchData = async () => {
      fetch(`${api_url}/api/scoreboard/team/${team}`, {
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
        {data && data.round >= 10 && (
          <DoubleArrowLeftIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href = `/scoreboard/team/${team}/${
                data.round - 10
              }`;
            }}
          />
        )}
        {data && data.round >= 1 && (
          <ArrowLeftIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href = `/scoreboard/team/${team}/${
                data.round - 1
              }`;
            }}
          />
        )}
        {data && (
          <Typography component='h1' variant='h5'>
            Round {data.round}
          </Typography>
        )}
        <ArrowLeftIcon
          sx={{
            visibility: "hidden",
          }}
        />
        <ArrowLeftIcon
          sx={{
            visibility: "hidden",
          }}
        />
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
                    window.location.href = `/scoreboard/check/${check.name}`;
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
