<script>
    import { onMount } from "svelte";
    import PatNav from "../../../../components/patNav.svelte"
  import PatientTab from "../../../../components/patient_tab.svelte";

    let profile = {}

   onMount(()=>{
        fetch("http://localhost:8000/patient_profile",{
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
        <PatNav />
    </div>
    <div class="col-span-8 flex flex-col items-center justify-center">
        <div class="flex items-center justify-center space-x-5 gap-y-10 pt-4">
            <div class="h-64 w-48 rounded-2xl flex flex-col space-y-2 items-center">
                <div class="flex h-52 w-full rounded-2xl">
                    <img src={`http://localhost:8000/${profile.profileImage}`}  alt="" srcset="" class="object-cover h-full w-full rounded-full">
                </div>     
                <div class="flex flex-col items-center">
                <p class="text-gray-700 font-semibold text-sm"> {profile.firstName} {profile.lastName}</p>
                </div>               
            </div>
            <div class="flex flex-col space-y-1">
                <p class="text-gray-700 font-semibold text-sm">First Name:  {profile.firstName}</p>
                <p class="text-gray-700 font-semibold text-sm">Last Name:  {profile.lastName}</p>
                <p class="text-gray-700 font-semibold text-sm">Surname:  {profile.surname}</p>
                <p class="text-gray-700 font-semibold text-sm">Email:  {profile.email}</p>
                <p class="text-gray-700 font-semibold text-sm">Contact:  {profile.contact}</p>
                <p class="text-gray-700 font-semibold text-sm">Address: {profile.address}</p>         
                <p class="text-gray-700 font-semibold text-sm">DOB: {profile.dob}</p>
            </div>        
        </div>
        <div class="flex items-center justify-center space-x-8 pt-7">      
            <a href="/patient_profile_update" class="bg-blue-700 py-3 px-10 text-white rounded-lg hover:bg-blue-500">Update Profile</a>        
        </div>
    </div>

    <div class="col-span-3">
        <PatientTab />
      </div>
</div>