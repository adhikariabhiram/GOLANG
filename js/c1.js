function swap(x) {
    let y=""
    for(let i=0;i<x.length;i++)
      {
        if(x.charAt(i)==x.charAt(i).toLowerCase())
        {
          y+=x.charAt(i).toUpperCase()
        }
        else if(x.charAt(i)==x.charAt(i).toUpperCase())
        {
          y+=x.charAt(i).toLowerCase()
        }
      }
    return y;
  }
  
  let sen="The Quick Brown Fox"
  console.log(swap(sen))
  