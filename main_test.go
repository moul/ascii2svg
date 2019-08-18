package main

import "os"

func Example() {
	os.Args = []string{"hello", "world"}
	main()
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <svg width="800px" height="600px" viewBox="0 0 800 600" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
	//   <g>
	//     <rect x="0" y="0" width="800" height="600" fill="black" stroke="white"></rect>
	//     <foreignObject x="15" y="5" width="770" height="590">
	//       <div xmlns="http://www.w3.org/1999/xhtml" style="width:800px; height:600px; overflow-y:auto; color: white; font-size: 17px;">
	//         <pre>world</pre>
	//       </div>
	//     </foreignObject>
	//   </g>
	// </svg>
}
