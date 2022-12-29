import { UserApi }from '../uitls/mockApi'
export default () => { 
    const test = () => {    
        const r = UserApi.findById({id: 123}) 
        return r.then((r) => { 
            console.log(r.data)
            return r 
        }).catch((e) => e)
    } 
    return {
        test
    }
}