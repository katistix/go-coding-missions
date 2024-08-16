### The Story: The Conveyor Belt Controller Challenge

---

At **FlexiManufacture Inc.**, the assembly line is the heart of production, and conveyor belts are the arteries that keep it flowing. Recently, the company upgraded their conveyor belt systems, which now come from two different manufacturers. These new belts use distinct communication protocols, complicating the job for the operators who need a unified way to control them.

Your team is tasked with a critical mission: develop a **Conveyor Belt Controller Interface**. This tool must manage both old and new conveyor belts, each using a unique protocol. It needs to start and stop the belts, adjust their speeds, and provide real-time status updates. With deadlines looming, your tool will become the linchpin in maintaining a smooth and efficient production line.

### Existing Protocols Documentation

#### **Protocol A: Binary Protocol**
- **Connection:** TCP/IP
- **Data Format:** Binary
- **Commands:**
  - **Start Conveyor:** `0x01`
  - **Stop Conveyor:** `0x02`
  - **Set Speed:** `0x03 [Speed Value]` (e.g., `0x03 0x05` for speed 5)
- **Responses:**
  - **Success:** `0x00`
  - **Error:** `0xFF [Error Code]`

#### **Protocol B: Text-Based Protocol**
- **Connection:** TCP/IP
- **Data Format:** Plain Text
- **Commands:**
  - **Start Conveyor:** `"START"`
  - **Stop Conveyor:** `"STOP"`
  - **Set Speed:** `"SPEED <value>"` (e.g., `"SPEED 5"`)
- **Responses:**
  - **Success:** `"OK"`
  - **Error:** `"ERROR <Error Message>"`

#### Protocol C: JSON-Based Protocol
- **Connection:** TCP/IP
- **Data Format:** JSON
- **Commands:**
  - **Start Conveyor:** `{"command": "START"}`
  - **Stop Conveyor:** `{"command": "STOP"}`
  - **Set Speed:** `{"command": "SPEED", "value": <speed>}` (e.g., `{"command": "SPEED", "value": 5}`)
- **Responses:**
  - **Success:** `{"status": "OK"}`
  - **Error:** `{"status": "ERROR", "message": "<Error Message>"}`

### Deliverables

1. **Conveyor Belt Controller Interface:** A CLI tool for controlling and monitoring conveyor belts using both protocols.
2. **User Documentation:** Guide for setup, configuration, and usage.
3. **Protocol Documentation:** Detailed command formats and expected responses for both protocols.
4. **Testing Suite:** Tests to ensure reliable operation and error handling.

