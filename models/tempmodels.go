package models

const (
    TITLE string = "GutterDropper 5000"
    // gutter page
    GUTTER_COST_FT float32 = 1.75 
    GUTTER_MARGIN float32 = 23.00
    LBS_FEET_6INCH float32 = 2.24
)


// temp functions
func TempGutterSelections() []GutterCoil {
    return []GutterCoil{
        {ID: 1, Size: "6 inches", Color: "White", OnHand: 654.45, PoundsFeetMulti: LBS_FEET_6INCH, DisplayInformation: "6\" White Gutter Coil", Cost: GUTTER_COST_FT, Margin: GUTTER_MARGIN, Selling: CostToSelling(GUTTER_COST_FT, GUTTER_MARGIN)},
        {ID: 2, Size: "6 inches", Color: "Black", OnHand: 45.55, PoundsFeetMulti: LBS_FEET_6INCH, DisplayInformation: "6\" Black Gutter Coil", Cost: GUTTER_COST_FT, Margin: GUTTER_MARGIN, Selling: CostToSelling(GUTTER_COST_FT, GUTTER_MARGIN), },
        {ID: 3, Size: "6 inches", Color: "Wicker", OnHand: 578.50, PoundsFeetMulti: LBS_FEET_6INCH, DisplayInformation: "6\" Wicker Gutter Coil", Cost: GUTTER_COST_FT, Margin: GUTTER_MARGIN, Selling: CostToSelling(GUTTER_COST_FT, GUTTER_MARGIN), },
        {ID: 4, Size: "6 inches", Color: "Pebblestone Clay", OnHand: 120.89, PoundsFeetMulti: LBS_FEET_6INCH, DisplayInformation: "6\" Pebblestone Clay Gutter Coil", Cost: GUTTER_COST_FT, Margin: GUTTER_MARGIN, Selling: CostToSelling(GUTTER_COST_FT, GUTTER_MARGIN), },
    }
}