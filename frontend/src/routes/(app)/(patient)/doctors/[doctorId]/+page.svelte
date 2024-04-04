<script>
    import { onMount } from "svelte";
    import PatNav from "../../../../../components/patNav.svelte";
    import {page} from "$app/stores"
    import {goto} from "$app/navigation"
    import PatientTab from "../../../../../components/patient_tab.svelte";

    
    let doctor = {}
    const id = $page.params.doctorId
    let purpose = "";
    let patientProfile = {}
    
    onMount(()=>{
        fetch(`http://localhost:8000/get_doctor/${id}`, {
        headers:{
          "content-type":"application/json",
          Authorization:`Bearer ${localStorage.getItem('hospital_token')}`
        }
      })
      .then(res=>res.json())
      .then(data=> {doctor=data; console.log(data)})
      .catch(err=>console.log(err))

      fetch("http://localhost:8000/patient_home",{
            headers:{
                "content-type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
            }
        })
        .then(res => res.json())
        .then(data => {patientProfile=data})
        .catch(err => console.log(err))      
    })

    function requestAppointment(id){
      fetch(`http://localhost:8000/request_appointment`, {
        method:"POST",
        headers:{
          "content-type":"application/json",
          Authorization:`Bearer ${localStorage.getItem('hospital_token')}`
        },
        body:JSON.stringify({
          "doctorId":id,
          "doctorProfile":doctor.tou,
          "purpose":purpose
        })
      })
      .then(res=>res.json())
      .then(data=> {
        console.log(data); 
        goto(`/${patientProfile.ID}/appointments`)})
      .catch(err=>console.log(err))
    }
</script>

<div class="bg-gray-200 h-screen grid grid-cols-12">
  <div class="col-span-1">
      <PatNav />
  </div>
  <div class="col-span-8 flex justify-center">

    <div class="flex items-center justify-center space-x-10 gap-y-10 pt-4">
        <div class=" h-64 w-48 rounded-2xl flex flex-col space-y-2 items-center">
            <div class="flex h-52  w-full rounded-2xl">
                <img src={`http://localhost:8000/${doctor.tou}`}  alt="" srcset="" class="object-cover h-full w-full rounded-full">
            </div>     
            <div class="flex flex-col items-center">
              <p class="text-gray-700 font-semibold text-sm"> Dr. {doctor.firstName}</p>
              <p class="text-gray-900 font-normal text-xs">{doctor.specialty}</p>
            </div>               
          </div>
          <form method="post" class="flex flex-col space-y-6" on:submit|preventDefault={()=> requestAppointment(doctor.ID)}>
            <input type="text" name="" id="" placeholder="Appointment Day" 
              class="border-2 border-gray-800 p-3 rounded-lg" bind:value={purpose} required />
            <button class="flex items-center justify-center p-3 rounded-lg bg-blue-500 text-white font-medium hover:bg-blue-400">Request Appointment</button>
          </form>
    </div>
  </div>
  <div class="col-span-3">
    <PatientTab />
  </div>  
</div>