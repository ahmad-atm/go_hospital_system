<script>
    import {onMount} from "svelte"
    import {goto} from "$app/navigation";
    import Icon from '@iconify/svelte';
  
    // let user = [];
    let error
    let page = "Home"
    onMount(() => {
      if(localStorage.getItem("hospital_token") == null){
        //goto("/login")
        return
      }    
      fetch('http://localhost:8000',{
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
                goto("/home")
                break;
            }
        })
        .catch(err => console.log(err))
    })
  </script>

<header class="text-gray-600 body-font">
  <div class="container mx-auto flex flex-wrap p-5 flex-col md:flex-row items-center">
    <a href="/home" class="flex title-font font-medium items-center text-gray-900 mb-4 md:mb-0">
      <Icon icon="fxemoji:hospital" />
      <span class="ml-3 text-xl">OMRI Hospital</span>
    </a>
    <nav class="md:ml-auto flex flex-wrap items-center text-base justify-center">
      <a href="/register_patient" class="mr-5 hover:text-gray-900">Register</a>
      <a href="/login" class="mr-5 hover:text-gray-900">Login</a>
    </nav>
  </div>
</header>

<section class="text-gray-600 body-font">
  <div class="container mx-auto flex px-5 py-24 md:flex-row flex-col items-center">
    <div class="lg:flex-grow md:w-1/2 lg:pr-24 md:pr-16 flex flex-col md:items-start md:text-left mb-16 md:mb-0 items-center text-center">
      <h1 class="title-font sm:text-4xl text-3xl mb-4 font-medium text-gray-900">Professional Care
        <br class="hidden lg:inline-block">For Your Family
      </h1>
      <p class="mb-8 leading-relaxed">Good Health is the state of mental, physical and social well being of and it does not just mean absense of disease</p>
      <div class="flex justify-center">
        <a href="/login" class="inline-flex text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded text-lg">Book An Appointment</a>
        <a href="/register_patient" class="ml-4 inline-flex text-gray-700 bg-gray-100 border-0 py-2 px-6 focus:outline-none hover:bg-gray-200 rounded text-lg">Register</a>
      </div>
    </div>
    <div class="lg:max-w-lg lg:w-full md:w-1/2 w-5/6">
      <img class="object-cover object-center rounded" alt="hero" src="Hospital/1.jpg">
    </div>
  </div>
</section>

  <!-- <div class="flex w-full justify-center space-x-5">
    <div class="flex justify-center items-center rounded-xl w-full h-72">
      <img src="Hospital/1.jpg" class="object-cover h-full rounded-xl w-full" alt="">
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
  </div> -->

  <!-- <div class="flex w-full justify-center space-x-5 mt-10">
    <div class="flex flex-col justify-center items-center bg-blue-100 rounded-lg w-full h-72">
      <p class="font-medium text-gray-800">Address: </p>
      <p class="font-medium text-gray-800 px-5">No.544 DORAYI BABBA U/LIMAN,</p>
      <p class="font-medium text-gray-800 px-5">BEHIND IMAM WALI GENERAL HOSPITAL KANO</p>
    </div> -->

    <!-- <div class="flex flex-col justify-center items-center bg-blue-100 w-full rounded-lg mr-5 p-5">
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
  </div> -->

  <!-- <div class="flex flex-col justify-center items-center">
    <div class="flex space-x-10 justify-center mt-10">
      <div class="rounded-full">
        <img src="Hospital/3.jpg" class="object-fit px-2 rounded-full" width="150" height="150" alt="">
      </div>
      <div class="rounded-full">
        <img src="Hospital/2.jpg" class="object-fit px-2 rounded-full" width="150" height="150" alt="">
      </div>
      <div class="rounded-full">
        <img src="Hospital/4.jpg" class="object-fit px-2 rounded-full" width="150" height="150" alt="">
      </div>
    </div> -->
  <!-- </div>   -->
