import { RouterProvider, createBrowserRouter } from "react-router-dom";
import App from "./App";

const Router = () => {

    const router = createBrowserRouter([
        {
          path: "/",
          element: <App />
        },
        {
          path: "about",
          element: <div>About</div>,
        },
      ]);

      return (
        <RouterProvider router={router} />

      )
}

export default Router