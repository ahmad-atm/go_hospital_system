<script>
    import { onMount } from "svelte";
    import DocNav from "../../../../components/docNav.svelte"
  import AppointmentsTab from "../../../../components/appointments_tab.svelte";

    let profile = {}
       onMount(()=>{
        fetch("http://localhost:8000/doctor_profile",{
            headers:{
              "Content-Type":"application/json",
              Authorization:`Bearer ${localStorage.getItem("hospital_token")}`
            }
        })
        .then(res => res.json())
        .then(data=> {console.log(data); profile=data})
        .catch(err => console.log(err))
    })
</script>


<div class="bg-gray-200 h-screen grid grid-cols-12">
    <div class="col-span-1">
        <DocNav />
    </div>
    <div class="col-span-8 flex flex-col justify-center items-center ">
        <div class="flex space-x-5 gap-y-10 pt-4">
                <div class="h-64 w-48 rounded-2xl flex flex-col space-y-2 items-center">
                    <div class="flex h-52  w-full">
                        <img src={`http://localhost:8000/${profile.tou}`}  alt="" srcset="" class="object-cover h-full w-full rounded-full">
                    </div>     
                    <div class="flex flex-col items-center">
                    <p class="text-gray-700 font-semibold text-sm">Dr. {profile.firstName} {profile.lastName}</p>
                    </div>               
                </div>
                <div class="flex flex-col space-y-1">
                    <p class="text-gray-700 font-semibold text-sm">First Name:  {profile.firstName}</p>
                    <p class="text-gray-700 font-semibold text-sm">Last Name:  {profile.lastName}</p>
                    <p class="text-gray-700 font-semibold text-sm">Surname:  {profile.surname}</p>
                    <p class="text-gray-700 font-semibold text-sm">Email:  {profile.email}</p>
                    <p class="text-gray-700 font-semibold text-sm">Contact:  {profile.contact}</p>
                    <p class="text-gray-700 font-semibold text-sm">Address: {profile.address}</p>
                    <p class="text-gray-700 font-semibold text-sm">Department: {profile.department}</p>
                    <p class="text-gray-700 font-semibold text-sm">Specialty: {profile.specialty}</p>
                    <p class="text-gray-700 font-semibold text-sm">DOB: {profile.dob}</p>
                </div>        
            </div>
            <div class="flex items-center justify-center space-x-8 pt-7">      
                <a href="/doctor_profile_update" class="bg-teal-700 py-3 px-10 text-white rounded-lg hover:bg-teal-500">Update Profile</a>        
            </div>        
    </div>
    <div class="col-span-3">
        <AppointmentsTab />
    </div>
 
</div>
