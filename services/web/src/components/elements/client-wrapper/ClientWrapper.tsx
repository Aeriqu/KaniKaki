import { FCProps } from "@/types/FCProps";
import { useEffect, useState } from "react";

/**
 * ClientWrapper wraps components to ensure it only runs when it's on the
 * client.
 * @param param0 FC inputs {children, ...props}
 */
export default function ClientWrapper({ children, ...props }: Readonly<FCProps>) {
  const [hasMounted, setHasMounted] = useState(false)

  useEffect(() => {
    setHasMounted(true);
  }, []);

  if(!hasMounted) {
    return null;
  }

  return <div {...props}>{children}</div>
}