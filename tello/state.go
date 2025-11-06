package tello

type TemperatureValue struct {
  low int
  high int
}

type MultiDirectionalValue struct {
  x int
  y int
  z int
}

type State struct {
  pitch int
  roll int
  yaw int
  altitude int
  height int
  time int
  baro int
  speed MultiDirectionalValue
  acceleration MultiDirectionalValue
}
