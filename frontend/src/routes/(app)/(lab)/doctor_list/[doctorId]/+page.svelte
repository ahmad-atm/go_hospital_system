<script>
    import { onMount } from "svelte";
    import LabNav from "../../../../../components/labNav.svelte";
    import {page} from "$app/stores";
  
    let doctorId = $page.params.doctorId;
    let doctor = {}
  
    onMount(async()=>{
          fetch(`http://localhost:8000/get_doctor/${doctorId}`,{
          headers:{
              "Content-type":"application/json",
              Authorization: `Bearer ${localStorage.getItem('hospital_token')}`
          }
          })
          .then(res=>res.json())
          .then(data=>{doctor = data; console.log(data)})
          .catch(err=>console.log(err))
      })
  
  
  </script>
  
  <div>
      <LabNav />
      <p class="flex justify-center text-lg text-gray-700 font-semibold border-b border-gray-400">Doctor Information</p>
      <div class="flex justify-center space-x-5 gap-y-10 pt-10">
          <div class="bg-gray-100 h-64 w-48 rounded-2xl flex flex-col space-y-2 items-center  shadow-xl">
              <div class="flex h-52 bg-white w-full rounded-2xl">
                  <img src={`http://localhost:8000/${doctor.tou}`}  alt="" srcset="" class="object-cover h-full w-full rounded-t-2xl">
              </div>     
              <div class="flex flex-col items-center">
                <p class="text-gray-700 font-semibold text-md">Dr {doctor.firstName}</p>
              </div>               
            </div>
          <div class="flex flex-col gap-y-2"> 
              <p class="text-gray-700 font-semibold text-md border-b border-gray-400">First Name:  {doctor.firstName}</p>
              <p class="text-gray-700 font-semibold text-md border-b border-gray-400">Last Name:  {doctor.lastName}</p>
              <p class="text-gray-700 font-semibold text-md border-b border-gray-400">Surname:  {doctor.surname}</p>
              <p class="text-gray-700 font-semibold text-md border-b border-gray-400">Email:  {doctor.email}</p>
              <p class="text-gray-700 font-semibold text-md border-b border-gray-400">Gender:  {doctor.gender}</p>
              <p class="text-gray-700 font-semibold text-md border-b border-gray-400">Contact:  {doctor.contact}</p>
              <p class="text-gray-700 font-semibold text-md border-b border-gray-400">Department:  {doctor.status}</p>
              <p class="text-gray-700 font-semibold text-md border-b border-gray-400">Specialty:  {doctor.specialty}</p>
              <p class="text-gray-700 font-semibold text-md border-b border-gray-400">Blood Type: {doctor.bloodType}</p>
              <p class="text-gray-700 font-semibold text-md border-b border-gray-400">Stop Time: {doctor.stop}</p>
  
          </div>        
      </div>
   
  </div>