import Head from 'next/head'
import useUserApi from '../../core/api/user/tmp' 
import useHostApi from '../../core/api/event/EventHostAPI' 
export default function Tmp() {
  const { createNewEvent } = useHostApi()
  return (
    <> 
      <button onClick={() => {
        console.log("clicked!") 
        createNewEvent(console.log, console.log)(
          { user_id: "abc" },
          {
            event_name: " string;",
            max_member: 1,
            detail: "string;",
            location: " string;"
          }
        )
      }} >
        click here to call mock API!!
      </button>
    </>
  )
}
