import axios from "axios";
import { useEffect, useState } from "react";

export function Shop() {
  const [products, setData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    axios.get('http://backend-api-endpoint/products')
      .then(response => {
        setData(response.data);
        setLoading(false);
      })
      .catch(error => {
        setError(error);
        setLoading(false);
      });
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  return (
    <main className="flex flex-col items-center justify-center pt-16 pb-4">
      <header className="flex items-center justify-center w-full">
        <div className="p-4">
          <h1 className="text-4xl">Flip Tech Test</h1>
        </div>
      </header>
      <div className="flex-1 flex items-center justify-center p-10 min-h-0">
        <table className="table-auto border-collapse border border-gray-200">
          <thead>
            <tr>
              <th className="border border-gray-300 px-4 py-2">SKU</th>
              <th className="border border-gray-300 px-4 py-2">Name</th>
              <th className="border border-gray-300 px-4 py-2">Price</th>
              <th className="border border-gray-300 px-4 py-2">Inventory Qty</th>
            </tr>
          </thead>
          <tbody>
            {products.map((item, index) => (
              <tr key={index}>
              <td className="border border-gray-300 px-4 py-2">{item.sku}</td>
              <td className="border border-gray-300 px-4 py-2">{item.name}</td>
              <td className="border border-gray-300 px-4 py-2">${item.price}</td>
              <td className="border border-gray-300 px-4 py-2">{item.inventory}</td>
            </tr>
            ))}
          </tbody>
        </table>
      </div>
    </main>
  )
}
