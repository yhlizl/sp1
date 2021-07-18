
 const url = location.protocol + '//' + location.host;

 async function senddata() {
     let result
     tempurl = url + "/sp1/healthy/get-config"
     let item = "sp1_ini"
     let fab = []
     document.querySelectorAll('#fab option:checked').forEach(v => {
         fab.push(v.value)
     })
     let phase = []
     document.querySelectorAll('#phase option:checked').forEach(v => {
         phase.push(v.value)
     })

    //  console.log(item, fab, phase)
     await axios.post(tempurl, {
             item: item,
             fab: fab,
             phase: phase,
         })
         .then((data) => {
             result = data
             gettable(result.data.data)
         })
         .catch(err => {
             console.log(err)
         })
     return result
 }

 function gettable(result) {
     let table = document.querySelector('#configtable')
     let tbody=table.querySelector("tbody")
     tbody.innerHTML=''
     let container = []
     console.log(result)

    // for(let i=0;i<Object.keys(result).length;i++){
    //     let temp={}
    //     console.log(result[i])
    //      temp["check"]=result[i]["check"]
    //      temp["params"]=result[i].params
    //      temp["value"]=result[i].value
    //      container.push(temp)
    // }

    container=Object.keys(result).map(key=>result[key])
    container.forEach(v =>{
        let tr= document.createElement("tr",)
        let td= document.createElement("td",)
        tr.insertBefore(td,null)
        Object.keys(v).forEach(key=>{
            console.log(v,key,v[key])
            let td= document.createElement("td",)
            let text =JSON.stringify(v[key])
            if (text =="true"){
                td.classList.add("bg-danger")
            }
            td.innerText=text
            tr.insertBefore(td,null)
        })
        tbody.insertBefore(tr,null)
    })
    table.DataTable()
     console.log(container)
 }

 function togglequery() {
     let getconfig = document.querySelector('#getconfig')
         //  console.log(getconfig.classList)
     if (getconfig.classList.contains("btn-primary")) {
         getconfig.classList.remove("btn-primary")
         getconfig.classList.add("btn-light")
         getconfig.setAttribute("disabled", "")
     } else {
         getconfig.classList.remove("btn-light")
         getconfig.classList.add("btn-primary")
         getconfig.removeAttribute("disabled")
     }
 }
 window.onload = function() {
     console.log("window loaded")

     let getconfig = document.querySelector('#getconfig')
     getconfig.addEventListener('click', async function(e) {
         e.preventDefault()
         togglequery()
         let data = await senddata()
         console.log(data)
         togglequery()
     })
 }