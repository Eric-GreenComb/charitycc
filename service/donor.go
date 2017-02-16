package service

import (
	"encoding/base64"
	"encoding/json"

	"github.com/CebEcloudTime/charitycc/core/store"
	"github.com/CebEcloudTime/charitycc/errors"
	"github.com/CebEcloudTime/charitycc/protos"
	"github.com/CebEcloudTime/charitycc/utils"
)

// RegisterDonor register donor
func RegisterDonor(store store.Store, args []string) ([]byte, error) {

	_donorAddr := args[1]
	_donorPublicKey := args[2]

	_, err := InitAccount(store, _donorAddr, _donorPublicKey)
	if err != nil {
		return nil, err
	}

	_, err = InitDonor(store, _donorAddr)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func InitDonor(store store.Store, addr string) ([]byte, error) {
	var newDonor protos.Donor
	newDonor.Addr = addr

	if tmpDonor, err := store.GetDonor(newDonor.Addr); err == nil && tmpDonor != nil && tmpDonor.Addr == newDonor.Addr {
		return nil, errors.AlreadyRegisterd
	}

	if err := store.PutDonor(&newDonor); err != nil {
		return nil, err
	}

	return nil, nil

}

// AddContribution add donor contribution
func AddContribution(store store.Store, args []string) ([]byte, error) {

	donorAddr := args[1]
	base64ContributionData := args[2]

	contributionData, err := base64.StdEncoding.DecodeString(base64ContributionData)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	newContribution, err := utils.ParseDonorContributionByJsonBytes(contributionData)
	if err != nil {
		return nil, err
	}

	tmpDonor, err := store.GetDonor(donorAddr)
	if err != nil {
		return nil, err
	}

	tmpDonor.Contributions = append(tmpDonor.Contributions, newContribution)

	if err := store.PutDonor(tmpDonor); err != nil {
		return nil, err
	}

	return nil, nil
}

// AddTrack add donor track
func AddTrack(store store.Store, args []string) ([]byte, error) {

	donorAddr := args[1]
	base64TrackData := args[2]

	trackData, err := base64.StdEncoding.DecodeString(base64TrackData)
	if err != nil {
		return nil, errors.Base64Decoding
	}

	newTracking, err := utils.ParseDonorTrackByJsonBytes(trackData)
	if err != nil {
		return nil, err
	}

	return PutTracking2Donor(store, donorAddr, newTracking)
}

// PutTracking2Donor
func PutTracking2Donor(store store.Store, donorAddr string, tracking *protos.DonorTrack) ([]byte, error) {
	tmpDonor, err := store.GetDonor(donorAddr)
	if err != nil {
		return nil, err
	}

	tmpDonor.Trackings = append(tmpDonor.Trackings, tracking)

	if err := store.PutDonor(tmpDonor); err != nil {
		return nil, err
	}

	return nil, nil
}

// QueryDonor get a donor
func QueryDonor(store store.Store, args []string) ([]byte, error) {

	addr := args[0]

	donor, err := store.GetDonor(addr)
	if err != nil {
		return nil, err
	}

	return json.Marshal(donor)
}

// QueryDonorObj
func QueryDonorObj(store store.Store, addr string) (*protos.Donor, error) {

	donor, err := store.GetDonor(addr)
	if err != nil {
		return nil, err
	}

	return donor, nil
}
