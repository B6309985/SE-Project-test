import React, { useEffect } from "react";
//import { useHistory } from "react-router-dom";

import { Link, Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { TreatmentsInterface } from "../models/ITreatment";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import axios from 'axios';
import { ButtonGroup } from "@mui/material";

function Treatment(props: any){
    const [treatment, setTreatment] = React.useState<TreatmentsInterface[]>([]);
    const getTreatment = async () => {
        const apiUrl = "http://localhost:8080/treatments";
        const requestOptions = {
            method: "GET",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json",
            },
        };

        await fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setTreatment(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    const handleDelete = async (id: number) => {
        try {
            const response = await axios.delete(`http://localhost:8080/treatments/${id}`, {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                    'Content-Type': 'application/json',
                }
            });

            if (response.status === 200) {
                console.log("Treatment deleted successfully");

                getTreatment();
            } else {
                throw new Error("Failed to delete treatment");
            }
        } catch (err) {
            console.error(err);
        }
    };

    const columns: GridColDef[] = [
        { field: "id", headerName: "ID", width: 70 },

        { field: "dentist_name", headerName: "Dentist", width: 70 },

        { field: "patient_name", headerName: "Patient", width: 70 },

        { field: "number_of_cavities", headerName: "Number of cavities", width: 130 },

        { field: "number_of_swollen_gums", headerName: "Number of swollen_gums", width: 150 },

        { field: "number_of_treatment", headerName: "number_of_treatment", width: 200 },

        { field: "other_teeth_problems", headerName: "Other teeth problems", width: 130 },

        { field: "type_of_treatment_name", headerName: "Type of treatment", width: 130 },

        

        

        { field: "treatment_detail", headerName: "Treatment", width: 90 },
 
        { field: "treatment_time", headerName: "Treatment Time", width: 120 },

        { field: "treatment_code", headerName: "Treatment code", width: 120 },

        // {
        //     field: "action", headerName: "Action",width: 200, sortable: false, renderCell: ({ row }) =>
        //     <ButtonGroup>
        //         <Button onClick={() => handleDelete(row.id)}>
        //             delete
        //         </Button>
        //         <Button >
        //             update
        //         </Button>
        //     </ButtonGroup>
        // },
        {
            field: "action", headerName: "Action",width: 200, sortable: false, renderCell: ({ row }) =>
            <ButtonGroup>
                <Button onClick={() => handleDelete(row.id)}>
                    delete
                </Button>
                <Button component={RouterLink} to={`/treatmentsupdate/${row.id}`} variant="contained">
                            <div className="good-font">
                                update
                            </div>
                        </Button>
            </ButtonGroup>
        },

    ];

    useEffect(() => {
        getTreatment();
    }, []);

    return (
        <div>
            <Container maxWidth="lg">
                <Box
                    display="flex"
                    sx={{
                        marginTop: 2,
                    }}
                >
                    <Box flexGrow={1}>
                        <Typography
                            component="h2"
                            variant="h6"
                            color="primary"
                            gutterBottom
                        >
                            <div className="good-font">
                                สาขา
                            </div>
                        </Typography>
                    </Box>
                    <Box>
                        <Button
                            component={RouterLink}
                            to="/treatmentscreate"
                            variant="contained"
                            color="primary"
                        >
                            <div className="good-font-white">
                                เพิ่มข้อการรักษา
                            </div>
                        </Button>
                    </Box>
                </Box>
                <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
                    <DataGrid
                        rows={treatment}
                        columns={columns}
                        pageSize={5}
                        rowsPerPageOptions={[5]}
                    />
                </div>
            </Container>
        </div>
    );
}

export default (Treatment);