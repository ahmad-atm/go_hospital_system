<script>
    import {goto} from "$app/navigation"
    import { onMount } from "svelte";
    import {page} from "$app/stores";
    import Select from "../../../../../components/select.svelte";
    import LabNav from "../../../../../components/labNav.svelte";

    let patient_id = $page.params.patient_id
    let patient = {}
    let doctors = []
    let doctor = {};

    let patientId = "";
    let patientFirstName = "";
    let patientLastName = ""
    let patientSurname = ""
    let patientImage = ""
    let status = "Pending"

    let tests = []
    let test = {}

    let title = "";
    let description = "";
    let price = ""
    onMount(async()=>{
        fetch(`http://localhost:8000/get_patient/${patient_id}`,{
        headers:{
            "Content-type":"application/json",
            Authorization: `Bearer ${localStorage.getItem('hospital_token')}`
        }
        })
        .then(res=>res.json())
        .then(data=>{
            patient = data; 
            patientId = patient.ID
            patientFirstName = patient.firstName;
            patientLastName = patient.lastName;
            patientSurname = patient.surname;
            patientImage = patient.profileImage;     
        })
        .catch(err=>console.log(err))

        await fetch("http://localhost:8000/doctor_list", {
            headers:{
                "Content-type":"application/json",
                Authorization:`Bearer ${localStorage.getItem('hospital_token')}`
            }
        })
        .then(res => res.json())
        .then(data => {
            console.log(data); 
            doctors=data;
        })
        .catch(err => console.log(err))

        await fetch("http://localhost:8000/all_tests", {
            headers:{
                "Content-type":"application/json",
                Authorization:`Bearer ${localStorage.getItem('hospital_token')}`
            }
        })
        .then(res => res.json())
        .then(data => {
            console.log(data); 
            tests=data;
        })
        .catch(err => console.log(err))
    })



    let error = "";    
    function createTest(){
        error = ''   
        fetch("http://localhost:8000/create_patient_test", {
            method:"POST",
            headers:{
                "Content-Type":"application/json",
                Authorization : `Bearer ${localStorage.getItem('hospital_token')}`
            },
            body:JSON.stringify({ 
                "title":test['title'], "description":test["description"], "price":test['price'],
                patientId, patientFirstName, patientLastName, patientSurname, 
                patientImage,  status,
                "doctorId": doctor["ID"], 
                "doctorFirstName": doctor["firstName"], 
                "doctorLastName": doctor["lastName"] , 
                "doctorSurname": doctor["surname"]
            })
        })
        .then(res=> res.json())
        .then(data=> {
            console.log(data)            
            goto("/lab_home")
        })
        .catch(err => console.log(err))
    }     
</script>

<LabNav />
<form on:submit|preventDefault={createTest}>
    <div class="gap-y-7 flex flex-col items-center justify-center min-h-screen">
        <div class="flex flex-col justify-center pb-5 px-10 gap-y-5 min-w-[30%] shadow-lg shadow-gray-400 rounded-lg">                         
            <p class="flex items-center justify-center pt-5 border-b-2 border-rose-300 text-lg text-gray-700">
                Create Test
            </p>           
            <div class="flex space-x-3 items-center">
                <label for="firstName">Patient First Name:</label>
                <input type="text" name="firstName" id="firstName" disabled placeholder="Patient First Name" 
                bind:value={patientFirstName} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />
            </div>                     
            <div class="flex space-x-3 items-center">
                <label for="lastName">Last Name:</label>
                <input type="text" name="lastName" id="lastName" disabled placeholder="Patient Last Name" 
                    bind:value={patientLastName} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />
            </div>

            <div class="flex space-x-3 items-center">
                <label for="surname">Surname:</label>
                <input type="text" name="surname" id="surname" disabled placeholder="Patient Surname" 
                    bind:value={patientSurname} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />
            </div>            
            <Select options={tests}
            display_func={o => o.title}
            bind:value={test}/>                     
            <div class="flex space-x-3 items-center">
                <label for="title">*Title:</label>
                <input type="text" name="title" id="title" required placeholder="Title" 
                    bind:value={test['title']} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />
            </div>
            <div class="flex space-x-3 items-center">
                <label for="description">*Description:</label>
                <input type="text" name="description" id="description" required placeholder="description" 
                    bind:value={test['description']} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />
            </div>   
            <div class="flex space-x-3 items-center">
                <label for="price">*Price:</label>
                <input type="price" name="price" id="price" required 
                    bind:value={test['price']} placeholder="Test Price" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />
            </div>

            <Select options={doctors}
            display_func={o => o.firstName + " " + o.lastName}
            bind:value={doctor}/>

            <div class="flex space-x-3 items-center">
                <label for="doctorFirstName">Doctor First Name:</label>
                <input type="text" name="doctorFirstName" id="doctorFirstName" required disabled bind:value={doctor["firstName"]} class="border border-slate-800 p-1">                                
            </div>            

            <div class="flex space-x-3 items-center">
                <label for="doctorLastName">Doctor Last Name:</label>
                <input type="text" name="doctorLastName" id="doctorLastName" required disabled 
                    bind:value={doctor["lastName"]} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />
            </div> 

            <div class="flex space-x-3 items-center">
                <label for="doctorSurname">Doctor Surname:</label>
                <input type="text" name="doctorSurname" id="doctorSurname" required disabled  
                    bind:value={doctor["surname"]} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />
            </div>             
         </div>
        <button type="submit" class="bg-blue-800 p-3 rounded-lg text-gray-200">Create Lab Test</button>
    </div>
</form>

<style>
    
</style>