<script>
    import {goto} from "$app/navigation"
    var email =  "";
    var password = "";

    function loginUser() { 
        fetch("http://localhost:8000/login", {
            method:"POST",
            headers:({"Content-Type":"application/json"}),
            body:JSON.stringify({
                email,password
            })
        })
        .then(res=> res.json())
        .then(data=> {
            console.log(data)
            localStorage.setItem("ecommerce_token", data)
            goto("/home")
        })
        .catch(err => console.log(err))
    }    
</script>

<div class="flex h-screen items-center justify-center bg-gray-50">
    <!-- <div class="w-[50%] flex  items-center justify-center bg-blue-500 border-4 rounded-lg border-white">
        <p>Login</p>
    </div> -->

    <div class="flex items-center justify-center">
        <form on:submit|preventDefault={loginUser}>
            
            <div class="gap-y-7 flex flex-col bg-white shadow-lg shadow-gray-300 p-10 rounded-2xl">
               <p class="flex justify-center items-center">Login</p>
                <div class="flex flex-col gap-y-5">
                    <div class="flex space-x-4 items-center px-5">
                        <label for="email" class="text-gray-700">Email:</label>
                        <input type="email" name="email" id="email" bind:value={email} 
                            placeholder="Email" class="p-2 mb-2 border-b border-gray-600 outline-none text-gray-800" 
                            autocomplete="name" required />                                
                    </div>
                    <div class="flex space-x-4 items-center">
                        <label for="password" class="text-gray-700">Password:</label>
                        <input type="password" name="password" id="password" 
                            autocomplete="current-password" bind:value={password} 
                            placeholder="Password" class="p-2 mb-2 border-b border-gray-600 outline-none text-gray-800"
                            required/>
                    </div>
                 </div>
                 <button type="submit" class="bg-blue-600 p-3 rounded-lg text-gray-200 hover:bg-blue-400">Login</button>
                 <a href="/register" class="flex items-center justify-center text-blue-500">Sign Up</a>
            </div>
        </form>
    </div>
</div>



<style>
    
</style>