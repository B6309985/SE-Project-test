import React from "react";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";

import Treatment from "./components/Treatment";

import TreatmentCreate from "./components/TreatmentCreate"; 

import Treatment_plan from "./components/Treatment_plan";

import TreatmentPlanCreate from "./components/TreatmentPlanCreate";

import TreatmentUpdate from "./components/TreatmentUpdate";

export default function App() { 

return (

  <Router>

   <div>
 
   <Navbar />

   <Routes>

       {/* <Route path="/" element={<Treatments />} /> */}
   
      <Route path="/treatmentscreate" element={<TreatmentCreate />} /> 

      * <Route path="/treatmentsupdate/:id" element={<TreatmentUpdate />} /> 

      <Route path="/treatment_grid" element={<Treatment />} />

      <Route path="/treatment_plan_grid" element={<Treatment_plan />} />

      <Route path="/treatment_plan_create" element={<TreatmentPlanCreate />} />

      {/* <Route path="/treatment_plan_grid" element={<trea />} /> */}

   </Routes>

   </div>

  </Router>

);

}