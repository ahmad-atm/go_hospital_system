<script>
    import DocNav from "../../../../../components/docNav.svelte";
    import { onMount } from "svelte";
    import {page} from "$app/stores"
    
      let test = [];
      let doctorTestId = $page.params.doctorTestId
      console.log(doctorTestId)
    
      onMount(()=>{
        fetch(`http://localhost:8000/get_patient_test/${doctorTestId}`, {
          headers:{
            "Content-type":"application/json",
            Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
          }
        })
        .then(res=>res.json())
        .then(data=>{test=data; console.log(test)})
        .catch(err=>console.log(err))
      })
    </script>
    
    <DocNav />
    <p class="flex justify-center text-lg text-gray-700 font-semibold border-b border-gray-400">Patient Test Information</p>
    <div class="flex  justify-center space-x-5 gap-y-10 pt-4">
        <div class="flex items-center justify-center">
            
        </div>
        <div class="bg-gray-100 h-64 w-48 rounded-2xl flex flex-col space-y-2 items-center  shadow-xl">
            <div class="flex h-52 bg-white w-full rounded-2xl">
                <img src={`http://localhost:8000/${test.patientImage}`}  alt="" srcset="" class="object-cover h-full w-full rounded-t-2xl">
            </div>     
            <div class="flex flex-col items-center">
              <p class="text-gray-700 font-semibold text-sm">{test.patientFirstName}</p>
              <p class="text-gray-900 font-normal text-xs">{test.patientLastName}</p>
            </div>               
          </div>
        <div>            
            <p class="text-gray-700 font-semibold text-lg">Patient Information</p>
            <hr class="pb-3" />
            <p class="text-gray-700 font-semibold text-md">First Name:  {test.patientFirstName}</p>
            <p class="text-gray-700 font-semibold text-md">Last Name:  {test.patientLastName}</p>
            <p class="text-gray-700 font-semibold text-md">Surname:  {test.patientSurname}</p>
            <p class="text-gray-700 font-semibold text-lg mt-7">Test Data</p>
            <hr class="pb-3" />
            <p class="text-rose-700 font-semibold text-md">Test Info:  {test.title}</p>
            <p class="text-rose-700 font-semibold text-md">Test Info:  {test.description}</p>
            <p class="text-gray-700 font-semibold text-md">Test Price: NGN {test.price}</p>
            <p class="text-gray-700 font-semibold text-md">Test Status: {test.status}</p>
            <div>
                <p class="text-gray-700 font-semibold text-lg mt-7">Doctor Information</p>
                <hr class="pb-3" />
                <p class="text-rose-700 font-semibold text-md">Test Info:  {test.doctorFirstName}</p>
                <p class="text-rose-700 font-semibold text-md">Test Info:  {test.doctorLastName}</p>
                <p class="text-rose-700 font-semibold text-md">Test Info:  {test.doctorSurname}</p>
            </div>
   
            <p class="text-gray-700 font-semibold text-lg mt-7">Lab Admin Information</p>
            <hr class="pb-3" />
            <p class="text-rose-700 font-semibold text-md">Test Info:  {test.labAdminFirstName}</p>
            <p class="text-rose-700 font-semibold text-md">Test Info:  {test.labAdminLastName}</p>
            <p class="text-gray-700 font-semibold text-md">Test Price: NGN {test.labAdminSurname}</p>
        </div>        
    </div>
    <!-- <div class="flex items-center justify-center space-x-8 pt-7">      
        <button class="bg-rose-700 py-3 px-10 text-white rounded-lg hover:bg-rose-500">Decline</button>  
        <button class="bg-teal-700 py-3 px-10 text-white rounded-lg hover:bg-teal-500">Accept</button>        
    </div> -->