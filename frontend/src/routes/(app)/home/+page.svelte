<script>
    import {onMount} from "svelte"
    import {goto} from "$app/navigation";
  
    // let user = [];
    let error
    let page = "Home"
    onMount(async () => {
      if(localStorage.getItem("hospital_token") == null){
        //goto("/login")
        return
      }
    
      await fetch('http://localhost:8000',{
            headers:{
              "Content-Type":"application/json",
              Authorization:`Bearer ${localStorage.getItem("hospital_token")}`
            }
          })
          .then(res => res.json())
          .then(data =>  {
            console.log(data)
            switch (data.role) {
              case "Lab":
                goto("/lab_home")
                break;
              case "Doctor":
                goto("/doctor_home")
                break;
              case "Patient":
                goto("/patient_home")
                break;
              default:
                break;
            }
        })
        .catch(err => console.log(err))
    })
  </script>


<div class="bg-blue-200 h-full">
  <div class="my-10">
    <nav class="bg-blue-500 flex items-center space-x-6 p-4 mt-5">
      <div class="bg-white p-3">
        <a href="/home" class="text-lg text-blue-800 bg-white rounded-lg font-medium mr-10">New Age Diabetes Clinic</a>
      </div>
      <a href="/home" class="text-md text-gray-100 font-medium">Home</a>
      <!-- <a href="/home" class="text-md text-gray-100 font-medium">Services</a> -->
      <a href="/register_patient" class="text-md text-gray-100 font-medium">Register</a>
      <a href="/login" class="text-md text-gray-100 font-medium">Login</a>
    </nav>
  </div>

  <div class="flex w-full justify-center space-x-5">
    <div class="flex justify-center items-center rounded-xl w-full h-72">
      <img src="Hospital/4.jpg" class="object-contain h-full rounded-xl w-full" alt="">
    </div>

    <div class="flex flex-col justify-center items-center bg-blue-100 w-full rounded-lg mr-5 p-5">
      <p class="font-medium text-gray-800 border-b-2">Our Services</p>
      <ul class="space-y-2">
        <li class="font-medium text-gray-800">General Surgery</li>
        <li class="font-medium text-gray-800">Diabetes Specialist</li>
        <li class="font-medium text-gray-800">Gynecology</li>
        <li class="font-medium text-gray-800">Dental</li>
        <li class="font-medium text-gray-800">Maternity</li>
        <li class="font-medium text-gray-800">Laboratory Services</li>        
      </ul>
    </div>
  </div>

  <div class="flex w-full justify-center space-x-5 mt-10">
    <div class="flex flex-col justify-center items-center bg-blue-100 rounded-lg w-full h-72">
      <p class="font-medium text-gray-800">Address: </p>
      <p class="font-medium text-gray-800 px-5">No.544 DORAYI BABBA U/LIMAN,</p>
      <p class="font-medium text-gray-800 px-5">BEHIND IMAM WALI GENERAL HOSPITAL KANO</p>
    </div>

    <div class="flex flex-col justify-center items-center bg-blue-100 w-full rounded-lg mr-5 p-5">
      <p class="font-medium text-gray-800 border-b-2">Contact Number</p>
      <ul class="space-y-2">
        <div class="flex space-x-4">
          <ul>
            <li class="font-medium text-gray-800">09080816792</li>
            <li class="font-medium text-gray-800">09080816107</li>            
          </ul>
          <ul>
            <li class="font-medium text-gray-800">09054161383</li>
            <li class="font-medium text-gray-800">07046105674</li>  
          </ul>
        </div>
      </ul>
    </div>
  </div>

  <div class="flex flex-col justify-center items-center">
    <div class="flex space-x-10 justify-center mt-10">
      <div class="rounded-full">
        <img src="Hospital/3.jpg" class="object-contain px-2 rounded-full" width="150" height="150" alt="">
      </div>
      <div class="rounded-full">
        <img src="Hospital/2.jpg" class="object-contain px-2 rounded-full" width="150" height="150" alt="">
      </div>
      <div class="rounded-full">
        <img src="Hospital/3.jpg" class="object-contain px-2 rounded-full" width="150" height="150" alt="">
      </div>
    </div>
  </div>  
</div>