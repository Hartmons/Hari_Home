package annalyn


func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake == true {
        return false
    } else {
        return true
    }
}


func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake == true || archerIsAwake == true || prisonerIsAwake == true{
        return true
    } else {
        return false
    }
}


func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if archerIsAwake == false && prisonerIsAwake == true {
        return true
    } else {
        return false
    }
}


func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if knightIsAwake == false && archerIsAwake == false && petDogIsPresent == true && prisonerIsAwake == false {
        return true
    }else if knightIsAwake == false && archerIsAwake == false && petDogIsPresent == true && prisonerIsAwake == true{
        return true
    } else if knightIsAwake == false && archerIsAwake == false && prisonerIsAwake == true && petDogIsPresent == false{ 
        return true
    }else if knightIsAwake == true && archerIsAwake == false && prisonerIsAwake == false && petDogIsPresent == true{
        return true
    }else if knightIsAwake == true && archerIsAwake == false && prisonerIsAwake == true && petDogIsPresent == true{
        return true
    }else {
        return false
    }
}
