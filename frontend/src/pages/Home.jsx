import React from "react"

import Navbar from "../component/navbar"

function HomePage() {
    return (
        <div>
            <Navbar/>
            <div style={{padding : "20%"}}>
                <h1 className="display-4 mb-4 font-weight-bold text-black">
                    Ini HOMEPAGE 
                </h1>
            </div>
        </div>
    )
}

export default HomePage