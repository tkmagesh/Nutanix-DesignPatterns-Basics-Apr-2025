

#  Modular Smart Home Automation System

You are building a **Modular Smart Home Automation System** that supports:

1. Different **device types** (like Lights, Fans, Thermostats, etc.)
2. **Grouping devices** (e.g., "All Lights on First Floor", "Bedroom Devices")
3. Allowing **custom behaviors** to be attached dynamically (like energy logging, remote control, etc.)
4. A **central dashboard** that reflects real-time device status updates.

---

## **Design Pattern Mapping**

1. **Factory Pattern** –  
   Dynamically create different types of smart devices (Light, Fan, Thermostat) from a `DeviceFactory` based on a configuration or input.

2. **Composite Pattern** –  
   Devices can be grouped into a `DeviceGroup`, which allows treating a single device and a group of devices uniformly for operations like `turnOn()`, `turnOff()`, or `getStatus()`.

3. **Decorator Pattern** –  
   Add behaviors like `EnergyMonitoring`, `RemoteControlEnabled`, or `VoiceControl` to devices at runtime without modifying the base class.

4. **Observer Pattern** –  
   The dashboard UI subscribes to device state changes. When a device is turned on/off or changes its state, all subscribed observers (like the dashboard or mobile app) are notified.

---

## **Use Case**

> A user installs a smart home system. She uses a factory to create individual smart lights and thermostats. She groups all bedroom devices using a composite pattern. She then adds energy monitoring to the thermostat using a decorator. Meanwhile, the mobile app dashboard observes all devices to reflect their real-time state.

---

