package midi

func GetMidiMap() map[uint8]KeyColor {
	return map[uint8]KeyColor{
		KICK:                {0, 0, 1, ALL_LEDS},
		SNARE_HEAD:          {1, 0, 0, SNARE_LEDS},
		SNARE_RIM:           {1, 0, 0, SNARE_LEDS},
		SNARE_X_STICK:       {1, 0, 0, SNARE_LEDS},
		TOM_1_HEAD:          {1, 0, 0, TOM_1_LEDS},
		TOM_1_RIM:           {1, 0, 0, TOM_1_LEDS},
		TOM_2_HEAD:          {1, 0, 0, TOM_2_LEDS},
		TOM_2_RIM:           {1, 0, 0, TOM_2_LEDS},
		TOM_3_HEAD:          {1, 0, 0, TOM_3_LEDS},
		TOM_3_RIM:           {1, 0, 0, TOM_3_LEDS},
		HIGHHAT_OPEN_BOW:    {0, 1, 0, SNARE_LEDS},
		HIGHHAT_OPEN_EDGE:   {0, 1, 0, SNARE_LEDS},
		HIGHHAT_CLOSED_BOW:  {0, 1, 0, SNARE_LEDS},
		HIGHHAT_CLOSED_EDGE: {0, 1, 0, SNARE_LEDS},
		HIGHHAT_PEDAL:       {0, 1, 0, SNARE_LEDS},
		CRASH_1_BOW:         {0, 1, 0, TOM_1_LEDS},
		CRASH_1_EDGE:        {0, 1, 0, TOM_1_LEDS},
		RIDE_BOW:            {0, 1, 0, TOM_2_LEDS},
		RIDE_EDGE:           {0, 1, 0, TOM_2_LEDS},
		RIDE_BELL:           {0, 1, 0, TOM_2_LEDS},
		CRASH_2_BOW:         {0, 1, 0, TOM_3_LEDS},
		CRASH_2_EDGE:        {0, 1, 0, TOM_3_LEDS},
	}
}
