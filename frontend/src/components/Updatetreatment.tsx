
import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import { useParams } from "react-router-dom";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import Autocomplete from "@mui/material/Autocomplete";

import { TreatmentsInterface } from "../models/ITreatment";
import { Type_of_treatments_Interface } from "../models/IType_of_treatment";
import { DentistsInterface } from "../models/IDentist";
import { PatientsInterface } from "../models/IPatient";
import { Type_of_number_of_treatment_Interface } from "../models/IType_of_number_of_treatment";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});


function TreatmentUpdate() {
    const [data, setData] = useState({});
    const [isLoading, setIsLoading] = useState(true);
    const [success, setSuccess] = React.useState(false);
    const [error, setError] = React.useState(false);
    const [errorMessage, setErrorMessage] = React.useState("");

    const [dentist, setADentist] = React.useState<DentistsInterface[]>([]); //React.useState<DentistsInterface>();
    const [patient, setPatient] = React.useState<PatientsInterface[]>([]);
    const [type_of_treatments, setType_of_treatments] = React.useState<Type_of_treatments_Interface[]>([]);
    const [type_of_number_of_treatments, setType_of_number_treatments] = React.useState<Type_of_number_of_treatment_Interface[]>([]);
    const [treatment, setTreatment] = React.useState<TreatmentsInterface>({ Treatment_time: new Date(), });
    //const [treatment_detail, setTreatment_detail] = React.useState<string>("");
    const [treatmentDetail, setTreatmentDetail] = React.useState("");
    const [treatment_code, setTreatment_code] = React.useState<string>("");

    const apiUrl = "http://localhost:8080";
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json"
        },
    };

    const { id } = useParams();

    const handleClose = (
        event?: React.SyntheticEvent | Event,
        reason?: string
    ) => {
        if (reason === "clickaway") {
            return;
        }
        setSuccess(false);
        setError(false);
    };

    const handleInputChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof TreatmentUpdate;
        const { value } = event.target;
        setTreatment({ ...treatment, [id]: value });
    };

    const handleChange = (event: SelectChangeEvent) => {
        const name = event.target.name as keyof typeof treatment;
        setTreatment({
            ...treatment,
            [name]: event.target.value,
        });
    };

    const getPatient = async () => {
        fetch(`${apiUrl}/patients`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setPatient(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    const getType_of_treatment = async () => {
        fetch(`${apiUrl}/type_of_treatments`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setType_of_treatments(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    const getType_of_number_of_treatment = async () => {
        fetch(`${apiUrl}/type_of_number_of_treatments`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setType_of_number_treatments(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    const getDentist = async () => {
        fetch(`${apiUrl}/dentists`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setADentist(res.data);
                }
                else { console.log("NO DATA") }
            });
    };


    // const value = parseInt(treatmentID!)
    // console.log("treatment by use Params id is"+treatmentID)

    useEffect(() => {
        //alert(id)
        getDentist();
        getPatient();
        getType_of_treatment();
        getType_of_number_of_treatment();
    }, []);

    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function submit() {
        let data = {
            DentistID: convertType(treatment.DentistID),
            PatientID: convertType(treatment.PatientID),

            Number_of_cavities: typeof treatment.number_of_cavities === "string" ? parseInt(treatment.number_of_cavities) : 0,
            Number_of_swollen_gums: typeof treatment.number_of_swollen_gums === "string" ? parseInt(treatment.number_of_swollen_gums) : 0,
            Other_teeth_problems: treatment.other_teeth_problems ?? "",
            Type_Of_TreatmentID: convertType(treatment.Type_of_treatmentsID),
            Number_of_treatment: typeof treatment.number_of_treatment === "string" ? parseInt(treatment.number_of_treatment) : 0,

            Type_Of_Number_Of_TreatmentID: convertType(treatment.Type_Of_Number_Of_TreatmentID),



            Treatment_detail: treatment.treatment_detail ?? "",
            Treatment_time: treatment.Treatment_time,
            Treatment_code: treatment.treatment_code ?? "",
        };

        console.log(data)

        useEffect(() => {
            async function fetchData() {
                const response = await fetch(`/treatments/${id}`);
                const json = await response.json();
                setData(json);
                setIsLoading(false);
            }
            fetchData();
        }, [id]);

        const [updatedData, setUpdatedData] = useState({});

        const handleChange = (event: { target: { name: any; value: any; }; }) => {
            const { name, value } = event.target;
            setUpdatedData({ ...updatedData, [name]: value });
        }

        const handleSubmit = async (event: { preventDefault: () => void; }) => {
            event.preventDefault();
            await fetch(`/treatments/${id}`, {
                method: 'PUT',
                body: JSON.stringify(updatedData),
                headers: { 'Content-Type': 'application/json' },
            });
        }

        if (isLoading) {
            return <p>Loading...</p>;
        }

        return (
            <form onSubmit={handleSubmit}>
                <input
                    name="field1"
                    value={data.Treatment_detail}
                    onChange={handleChange}
                />
                {/* <input
                name="field2"
                value={updatedData.field2 || data.field2}
                onChange={handleChange}
            />
            <button type="submit">Update</button> */}
            </form>
        );

    }
}

export default TreatmentUpdate;