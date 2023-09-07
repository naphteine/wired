import { createClient } from "@supabase/supabase-js";

export const supabase = createClient("https://byjylmkxlxicveuczdpw.supabase.co", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImJ5anlsbWt4bHhpY3ZldWN6ZHB3Iiwicm9sZSI6ImFub24iLCJpYXQiOjE2OTQxMDQ1MTQsImV4cCI6MjAwOTY4MDUxNH0.aQTwA4fjqV0v3iX-3kssVGx2DpL8VwbDc2A4eiQkI0I");